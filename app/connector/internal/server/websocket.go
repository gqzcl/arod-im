package server

import (
	"arod-im/app/connector/internal/conf"
	"arod-im/app/connector/internal/service"
	"arod-im/pkg/transport/websocket"
)

// NewWebsocketServer create a websocket server.
func NewWebsocketServer(c *conf.Server, s *service.ConnectorService) *websocket.Server {
	srv := websocket.NewServer(
		websocket.Address(c.Websocket.Addr),
		websocket.Path(c.Websocket.Path),
		websocket.OnMessageHandle(s.OnMessageHandler),
		websocket.OnCloseHandle(s.OnCloseHandler),
	)

	s.SetWebsocketServer(srv)

	return srv
}

//type wsServer struct {
//	gnet.BuiltinEventEngine
//	eng       gnet.Engine
//	multicore bool
//
//	//buckets   []*connector.Bucket
//	//rpcClient logic.LogicClient
//
//	addr      string
//	connected int64
//	serverID  string
//	i         bool
//}
//
//type wsCodec struct {
//	ws bool
//}
//
//// Onboot 启动时触发
//func (wss *wsServer) Onboot(eng gnet.Engine) gnet.Action {
//	wss.eng = eng
//	//ch := connector.NewChannel(nil)
//	//wss.buckets[0].AddChannel("test", ch)
//	logging.Infof("arod-im connector server with multi-core=%t is listening on %s", wss.multicore, wss.addr)
//	return gnet.None
//}
//
//// OnOpen 新连接打开时触发
//func (wss *wsServer) OnOpen(c gnet.Conn) ([]byte, gnet.Action) {
//	c.SetContext(new(wsCodec))
//
//	atomic.AddInt64(&wss.connected, 1)
//	//wss.connectedSockets.Set(c.RemoteAddr(), c)
//	return nil, gnet.None
//}
//
//// OnClose 连接关闭时触发，通知客户端断开连接，同时清除用户在线状态
//func (wss *wsServer) OnClose(c gnet.Conn, err error) (action gnet.Action) {
//	if err != nil {
//		logging.Warnf("error occurred on connection=%s, %v\n", c.RemoteAddr().String(), err)
//	}
//	// TODO 通知logic清除用户在线状态
//	atomic.AddInt64(&wss.connected, -1)
//	logging.Infof("conn[%v] disconnected", c.RemoteAddr().String())
//	return gnet.None
//}
//
//// OnTraffic 收到数据时触发，只处理心跳包，其他消息一律发送给logic
//func (wss *wsServer) OnTraffic(c gnet.Conn) gnet.Action {
//	if !c.Context().(*wsCodec).ws {
//		_, err := ws.Upgrade(c)
//		logging.Infof("conn[%v] upgrade websocket protocol", c.RemoteAddr().String())
//		if err != nil {
//			logging.Infof("conn[%v] [err=%v]", c.RemoteAddr().String(), err.Error())
//			return gnet.Close
//		}
//		c.Context().(*wsCodec).ws = true
//	} else {
//		// TODO auth，验证token
//		// TODO 如果是为鉴权连接则直接关闭
//		// TODO store conn
//		msg, op, err := wsutil.ReadClientData(c)
//		if err != nil {
//			if _, ok := err.(wsutil.ClosedError); !ok {
//				logging.Infof("conn[%v] [err=%v]", c.RemoteAddr().String(), err.Error())
//			}
//			return gnet.Close
//		}
//		logging.Infof("conn[%v] receive [op=%v] [msg=%v]", c.RemoteAddr().String(), op, string(msg))
//		// 处理心跳包
//		if int8(msg[0]) == protocol.OpHeartbeat {
//			err = wsutil.WriteServerMessage(c, ws.OpPong, []byte("heatbeat received"))
//			logging.Infof("conn[%v] [push=%s]", c.RemoteAddr().String(), "heatbeat received")
//			if err != nil {
//				logging.Infof("conn[%v] [err=%v]", c.RemoteAddr().String(), err.Error())
//				return gnet.Close
//			}
//		} else if int8(msg[0]) == protocol.OpAuth {
//			//TODO 鉴权
//		} else {
//			//将数据发到logic
//			_, err = wss.rpcClient.PushMsg(context.Background(), &logic.Msg{Msg: msg})
//			if err == nil {
//				logging.Infof("push to logic success")
//			} else {
//				logging.Infof("push to logic failed,error+%v", err)
//			}
//		}
//		fmt.Println(op, "string+", string(msg), "bin+", msg)
//		logging.Infof("[connected-count=%v]", atomic.LoadInt64(&wss.connected))
//	}
//	return gnet.None
//}
//
//// OnTick 定时任务
//func (wss *wsServer) OnTick() (delay time.Duration, action gnet.Action) {
//	// TODO 定时发送心跳包检测连接存活状况
//	if wss.i {
//		to := wss.buckets[0].GetChannel("test")
//		if to != nil {
//			err := to.Push([]byte("connect sucess"))
//			if err != nil {
//				logging.Infof("push to client failed,error+%v", err)
//			} else {
//				logging.Infof("push to client success")
//			}
//		}
//
//	}
//
//	//logging.Infof("[connected-count=%v]", atomic.LoadInt64(&wss.connected))
//	return 3 * time.Second, gnet.None
//}
//
//func (wss *wsServer) StartWebsocket(addrs []string) (err error) {
//
//	for _, bind := range addrs {
//		addr := fmt.Sprintf("tcp://127.0.0.1:%s", bind)
//		log.Infof("start wss listen on: ", bind)
//		//gnet.Run(wss, addr,
//		//	gnet.WithLoadBalancing(gnet.RoundRobin),
//		//	gnet.WithTCPKeepAlive(5),
//		//)
//		//log.Infoln("server exits:", gnet.Run(wss, addr, gnet.WithMulticore(true), gnet.WithReusePort(true), gnet.WithTicker(true)))
//		gnet.Run(wss, addr, gnet.WithMulticore(true), gnet.WithReusePort(true), gnet.WithTicker(true))
//	}
//	return
//}
