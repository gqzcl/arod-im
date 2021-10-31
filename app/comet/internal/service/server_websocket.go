package service

import (
	"context"
	"crypto/tls"
	"io"
	"net"
	"runtime"
	"strings"
	"time"

	"arod-im/api/protocol"
	"arod-im/app/comet/internal/conf"
	"arod-im/app/comet/internal/pkg"
	"arod-im/pkg/wbyte"
	"arod-im/pkg/websocket"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/golang/glog"
)

// InitWebsocket listen all tcp.bind and start accept connections.
func InitWebsocketServer(server *Server, addrs string, logger log.Logger) (err error) {
	log := log.NewHelper(log.With(logger, "module", "comet/tcpserver"))
	log.Info("start tcp server")
	addr, err := net.ResolveTCPAddr("tcp", addrs)
	if err != nil {
		log.Errorf("net.ResolveTCPAddr error:%s", err)
		return err
	}
	// part1: create a listener
	listen, err := net.ListenTCP("tcp", addr)
	if err != nil {
		log.Errorf("net.ListenTCP error:%s", err)
		return err
	}
	log.Infof("start ws listen: %s", addrs)
	for i := 0; i < runtime.NumCPU(); i++ {
		go acceptWebsocket(server, listen, log)
	}
	return nil
}

// InitWebsocketWithTLS init websocket with tls.
func InitWebsocketWithTLS(server *Server, addrs []string, certFile, privateFile string, logger log.Logger) (err error) {
	log := log.NewHelper(log.With(logger, "module", "comet/tcpserver"))
	var (
		bind     string
		listener net.Listener
		cert     tls.Certificate
		certs    []tls.Certificate
	)

	// tls 证书验证
	certFiles := strings.Split(certFile, ",")
	privateFiles := strings.Split(privateFile, ",")
	for i := range certFiles {
		cert, err = tls.LoadX509KeyPair(certFiles[i], privateFiles[i])
		if err != nil {
			glog.Errorf("Error loading certificate. error(%v)", err)
			return
		}
		certs = append(certs, cert)
	}
	tlsCfg := &tls.Config{Certificates: certs}

	tlsCfg.BuildNameToCertificate()
	// 为每个地址创建goroutine
	for _, bind = range addrs {
		if listener, err = tls.Listen("tcp", bind, tlsCfg); err != nil {
			glog.Errorf("net.ListenTCP(tcp, %s) error(%v)", bind, err)
			return
		}
		glog.Infof("start wss listen: %s", bind)
		// split N core accept
		for i := 0; i < runtime.NumCPU(); i++ {
			go acceptWebsocketWithTLS(server, listener, log)
		}
	}
	return
}

// Accept accepts connections on the listener and serves requests
// for each incoming connection.  Accept blocks; the caller typically
// invokes it in a go statement.
func acceptWebsocket(server *Server, listen *net.TCPListener, log *log.Helper) {
	var (
		conn *net.TCPConn
		r    int
		t    *conf.TCP
	)
	for {
		c, err := listen.AcceptTCP()
		if err != nil {
			log.Errorf("listen.Accept() error(%v)", err)
			return
		}
		defer c.Close()
		c.SetKeepAlive(t.KeepAlive)
		c.SetReadBuffer(int(t.Rcvbuf))
		c.SetWriteBuffer(int(t.Sndbuf))
		go handleWebsocket(server, conn, r, log)
		if r++; r == maxInt {
			r = 0
		}
	}
}

// Accept accepts connections on the listener and serves requests
// for each incoming connection.  Accept blocks; the caller typically
// invokes it in a go statement.
func acceptWebsocketWithTLS(server *Server, listen net.Listener, log *log.Helper) {
	var (
		conn net.Conn
		err  error
		r    int
	)
	for {
		if conn, err = listen.Accept(); err != nil {
			// if listener close then return
			glog.Errorf("listener.Accept(\"%s\") error(%v)", listen.Addr().String(), err)
			return
		}
		go handleWebsocket(server, conn, r, log)
		if r++; r == maxInt {
			r = 0
		}
	}
}

func handleWebsocket(server *Server, c net.Conn, r int, log *log.Helper) {
	var (
		round = &pkg.Round{}
		proto = &conf.Protocol{}
		times = round.Timer(r) // 倒计时
		rp    = round.Reader(r)
		wp    = round.Writer(r)
		// localAddr  = c.LocalAddr().String()
		remoteAddr = c.RemoteAddr().String()
		p          *protocol.Proto
		ch         = pkg.NewChannel(int(proto.CliProto), int(proto.SvrProto))
		rid        string
		accepts    []int32
		hb         time.Duration
		chread     = &ch.Reader
		chwrite    = &ch.Writer
		readbuf    = rp.Get()
		writebuf   = wp.Get()
		lasthb     = time.Now()
		ws         *websocket.Conn
		req        *websocket.Request
		b          *pkg.Bucket
		err        error
	)
	// 设置Channel的读缓冲区
	ch.Reader.ResetBuffer(c, readbuf.Bytes())

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	// Step 0
	// 握手
	timerdata := times.Add(proto.HandshakeTimeout.AsDuration(), func() {
		c.SetDeadline(time.Now().Add(time.Millisecond * 100))
		c.Close()
		log.Errorf("%s handshake timeout", remoteAddr)
	})
	ch.IP, _, _ = net.SplitHostPort(remoteAddr)

	// Step 1
	// 判断连接的url是否正确
	if req, err := websocket.ReadRequest(chread); err != nil || req.RequestURI != "/sub" {
		c.Close()
		times.Del(timerdata)
		rp.Put(readbuf)
		if err != io.EOF {
			glog.Errorf("http.ReadRequest(rr) error(%v)", err)
		}
		return
	}
	// 设置channel的写buf
	ch.Writer.ResetBuffer(c, writebuf.Bytes())
	// Step 2
	// 将TCP连接升级为Websocket
	if ws, err = websocket.Upgrade(c, chread, chwrite, req); err != nil {
		// 关闭连接
		c.Close()
		// 删除心跳计时
		times.Del(timerdata)
		// 回收读写buf
		rp.Put(readbuf)
		wp.Put(writebuf)
		if err != io.EOF {
			log.Errorf("websocket.NewServerConn error(%v)", err)
		}
		return
	}
	// Step 3
	// websocket auth
	if p, err := ch.CliProto.Write(); err == nil {
		if ch.MemberId, ch.Key, rid, accepts, hb, err = authWebsocket(ctx, server, ws, p, req.Header.Get("Cokkie")); err == nil {
			ch.Watch(accepts...)
			// 将user channel放到bucket中
			b = server.GetBucket(ch.Key)
			err = b.PutChannel(rid, ch)
		}
	}
	// Step 4
	// 如果有错误
	if err != nil {
		// 关闭连接
		ws.Close()
		// 回收读写buf
		rp.Put(readbuf)
		wp.Put(writebuf)
		// 移除心跳计时
		times.Del(timerdata)
		return
	}
	// 成功后重置心跳时间
	timerdata.Key = ch.Key
	times.Set(timerdata, hb)

	// Step 5
	// 分发TCP请求
	go dispatchWebsocket(ws, wp, writebuf, ch)
	// 设置心跳时间
	serverHeartbeat := server.RandServerHearbeat()
	// 处理TCP数据
	for {
		// 读取TCP错误，退出
		if p, err = ch.CliProto.Read(); err != nil {
			break
		}
		if err = p.ReadWebsocket(ws); err != nil {
			break
		}
		// 回应心跳
		if p.Op == protocol.OpHeartbeat {
			// comet有心跳機制維護連線狀態，對於logic來說也需要有人利用心跳機制去告知哪個user還在線
			// 目前在不在線這個狀態都是由comet控管，但不需要每次tcp -> 心跳 -> comet就 -> 心跳 -> logic
			// 所以webSocket -> comet 心跳週期會比 comet -> logic還要短
			// 假設
			// 1. tcp -> comet 5分鐘沒心跳就過期
			// 2. comet -> logic 20分鐘沒心跳就過期
			// tcp -> 每30秒心跳 -> comet <====== 每次只要不超過5分鐘沒心跳則comet會認為連線沒問題
			// tcp -> 每30秒心跳 -> comet -> 判斷是否已經快20分鐘沒通知logic(是就發) -> logic
			times.Set(timerdata, hb)
			p.Op = protocol.OpHeartbeatReply
			p.Body = nil
			if now := time.Now(); now.Sub(lasthb) > serverHeartbeat {
				if err = server.Heartbeat(ctx, ch.MemberId, ch.Key); err == nil {
					lasthb = now
				}
			}
		} else {
			// 处理其他操作
			if err = server.Operate(ctx, p, ch, b); err != nil {
				break
			}
		}
		// 写入指针向前移动
		ch.CliProto.WriteAdv()
		// 通知goroutine处理本次收到的请求
		ch.Signal()
	}
	// 连接异常时需要从bucket中将channel移除
	b.DelChannel(ch)
	// 移除心跳任务
	times.Del(timerdata)
	// 回收读buf
	rp.Put(readbuf)
	// 关闭连接
	ws.Close()
	// 发送结束信号关闭channel
	ch.Close()
	// 通知Logic当前用户下线
	if err = server.Disconnect(ctx, ch.MemberId, ch.Key); err != nil {
		log.Errorf("key: %s disconnect error(%v)", ch.Key, err)
	}
}

func dispatchWebsocket(ws *websocket.Conn, wp *wbyte.Pool, wb *wbyte.Buffer, ch *pkg.Channel) {
	var (
		err    error
		finish bool
		online int32
	)
	for {
		// 等待接收通知，没有通知时阻塞
		var p = ch.Ready()
		switch p {
		// 收到结束信号时进入failed，处理连接关闭
		case protocol.ProtoFinish:
			finish = true
			goto failed
		// 收到Ready信号时开始处理信息
		case protocol.ProtoReady:
			for {
				// 从Ringbuf中读取proto
				if p, err = ch.CliProto.Read(); err != nil {
					break
				}
				// 如果是心跳操作，写入心跳信息
				if p.Op == protocol.OpHeartbeatReply {
					if ch.Room != nil {
						online = ch.Room.OnlineNum()
					}
					if err = p.WriteWebsocketHeart(ws, online); err != nil {
						goto failed
					}
				} else {
					// 如果是其他操作则将proto写入
					if err = p.WriteWebsocket(ws); err != nil {
						goto failed
					}
				}
				p.Body = nil // avoid memory leak
				// 读指针前进
				ch.CliProto.ReadAdv()
			}
		default:
			// server send
			if err = p.WriteWebsocket(ws); err != nil {
				goto failed
			}

		}
		// only hungry flush response
		if err = ws.Flush(); err != nil {
			break
		}

	}
	// 关闭连接
	// 回收写buff
failed:
	ws.Close()
	wp.Put(wb)
	// must ensure all channel message discard, for reader won't blocking Signal
	for !finish {
		finish = (ch.Ready() == protocol.ProtoFinish)
	}
}

// auth for goim handshake with client, use rsa & aes.
func authWebsocket(ctx context.Context, server *Server, ws *websocket.Conn, p *protocol.Proto, cookie string) (mid int64, key, rid string, accepts []int32, hb time.Duration, err error) {
	for {
		if err = p.ReadWebsocket(ws); err != nil {
			return
		}
		if p.Op == protocol.OpAuth {
			break
		} else {
			glog.Errorf("ws request operation(%d) not auth", p.Op)
		}
	}
	if mid, key, rid, accepts, hb, err = server.Connect(ctx, p, cookie); err != nil {
		return
	}
	p.Op = protocol.OpAuthReply
	p.Body = nil
	if err = p.WriteWebsocket(ws); err != nil {
		return
	}
	err = ws.Flush()
	return
}
