package service

import (
	"arod-im/api/protocol"
	"arod-im/app/comet/internal/conf"
	"arod-im/app/comet/internal/pkg"
	"arod-im/pkg/selfbufio"
	"arod-im/pkg/wbyte"
	"context"
	"net"
	"runtime"
	"time"

	"github.com/go-kratos/kratos/v2/log"
)

func InitTCPServer(server *Server, addrs string, logger log.Logger) error {
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
		go acceptTCP(server, listen, log)
	}
	return nil
}

const maxInt = 1<<31 - 1

func acceptTCP(server *Server, listen *net.TCPListener, log *log.Helper) {
	var r int
	var t *conf.TCP
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
		go handleTCP(server, c, r, log)
		if r++; r == maxInt {
			r = 0
		}
	}
}
func handleTCP(server *Server, c *net.TCPConn, r int, log *log.Helper) {
	var (
		round = &pkg.Round{}
		proto = &conf.Protocol{}
		times = round.Timer(r) // 倒计时
		rp    = round.Reader(r)
		wp    = round.Writer(r)
		// localAddr  = c.LocalAddr().String()
		remoteAddr = c.RemoteAddr().String()
		ch         = pkg.NewChannel(int(proto.CliProto), int(proto.SvrProto))
		rid        string
		accepts    []int32
		hb         time.Duration
		chread     = &ch.Reader
		chwrite    = &ch.Writer
		readbuf    = rp.Get()
		writebuf   = wp.Get()
		lasthb     = time.Now()
	)
	// 设置Channel的读写缓冲区
	ch.Reader.ResetBuffer(c, readbuf.Bytes())
	ch.Writer.ResetBuffer(c, writebuf.Bytes())

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// 握手
	timerdata := times.Add(proto.HandshakeTimeout.AsDuration(), func() {
		c.Close()
		log.Errorf("%s handshake timeout", remoteAddr)
	})
	ch.IP, _, _ = net.SplitHostPort(remoteAddr)

	// 鉴权
	p, err := ch.CliProto.Write()
	ch.MemberId, ch.Key, rid, accepts, hb, err = authTCP(ctx, chread, chwrite, p, server, log)
	ch.Watch(accepts...)
	b := server.GetBucket(ch.Key)
	err = b.PutChannel(rid, ch)
	log.Infof("tcp connnected key:%s mid:%d proto:%+v", ch.Key, ch.MemberId, p)

	// 失败时回收读写buf，删除心跳计时器并关闭连接
	if err != nil {
		c.Close()
		rp.Put(readbuf)
		wp.Put(writebuf)
		times.Del(timerdata)
		log.Errorf("key: %s handshake failed error(%v)", ch.Key, err)
		return
	}
	// 成功后重置心跳时间
	timerdata.Key = ch.Key
	times.Set(timerdata, hb)

	// 分发TCP请求
	go dispatchTCP(c, chwrite, wp, writebuf, ch)
	// 设置心跳时间
	serverHeartbeat := server.RandServerHearbeat()
	// 处理TCP数据
	for {
		// 读取TCP错误，退出
		if p, err = ch.CliProto.Read(); err != nil {
			break
		}
		if err = p.ReadTCP(chread); err != nil {
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
	c.Close()
	// 发送结束信号关闭channel
	ch.Close()
	// 通知Logic当前用户下线
	if err = server.Disconnect(ctx, ch.MemberId, ch.Key); err != nil {
		log.Errorf("key: %s disconnect error(%v)", ch.Key, err)
	}
}

func authTCP(ctx context.Context, rr *selfbufio.Reader, wr *selfbufio.Writer, p *protocol.Proto, server *Server, log *log.Helper) (mid int64, key, rid string, accepts []int32, hb time.Duration, err error) {
	for {
		if err = p.ReadTCP(rr); err != nil {
			return
		}
		if p.Op == protocol.OpAuth {
			break
		} else {
			log.Errorf("tcp request operation(%d) not auth", p.Op)
		}
	}
	// 连接
	if mid, key, rid, accepts, hb, err = server.Connect(ctx, p, ""); err != nil {
		log.Errorf("authTCP.Connect(key:%v).err(%v)", key, err)
		return
	}
	// 回复
	p.Op = protocol.OpAuthReply
	p.Body = nil
	if err = p.WriteTCP(wr); err != nil {
		log.Errorf("authTCP.WriteTCP(key:%v).err(%v)", key, err)
		return
	}
	err = wr.Flush()
	return
}

func dispatchTCP(conn *net.TCPConn, wr *selfbufio.Writer, wp *wbyte.Pool, wb *wbyte.Buffer, ch *pkg.Channel) {
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
					if err = p.WriteTCPHeart(wr, online); err != nil {
						goto failed
					}
				} else {
					// 如果是其他操作则将proto写入
					if err = p.WriteTCP(wr); err != nil {
						goto failed
					}
				}
				p.Body = nil // avoid memory leak
				// 读指针前进
				ch.CliProto.ReadAdv()
			}
		default:
			// server send
			if err = p.WriteTCP(wr); err != nil {
				goto failed
			}

		}
		// only hungry flush response
		if err = wr.Flush(); err != nil {
			break
		}

	}
failed:

	conn.Close()
	wp.Put(wb)
	// must ensure all channel message discard, for reader won't blocking Signal
	for !finish {
		finish = (ch.Ready() == protocol.ProtoFinish)
	}

}
