// Copyright 2022 gqzcl <gqzcl@qq.com>. All rights reserved.
// Use of this source code is governed by a MIT style

package websocket

import (
	"sync/atomic"
	"time"

	"github.com/gobwas/ws/wsutil"
	"github.com/panjf2000/gnet/v2"
)

func (s *Server) OnBoot(eng gnet.Engine) gnet.Action {
	s.eng = eng
	s.log.Infof("ws server with multi-core=%t is listening on %s", true, s.address)
	return gnet.None
}

func (s *Server) OnOpen(c gnet.Conn) ([]byte, gnet.Action) {
	c.SetContext(new(WsContext))
	atomic.AddInt64(&s.connected, 1)
	return nil, gnet.None
}

func (s *Server) OnClose(c gnet.Conn, err error) (action gnet.Action) {
	if err != nil {
		s.log.Warnf("error occurred on connection=%s, %v\n", c.RemoteAddr().String(), err)
	}
	atomic.AddInt64(&s.connected, -1)
	s.onCloseHandler(c)
	s.log.Infof("conn[%v] disconnected", c.RemoteAddr().String())
	return gnet.None
}

func (s *Server) OnTraffic(c gnet.Conn) gnet.Action {
	if !c.Context().(*WsContext).ws {
		_, err := s.upgrader.Upgrade(c)
		s.log.Infof("conn[%v] upgrade websocket protocol", c.RemoteAddr().String())
		if err != nil {
			s.log.Infof("conn[%v] [err=%v]", c.RemoteAddr().String(), err.Error())
			return gnet.Close
		}
		c.Context().(*WsContext).ws = true
	} else {
		msg, op, err := wsutil.ReadClientData(c)
		if err != nil {
			if _, ok := err.(wsutil.ClosedError); !ok {
				s.log.Infof("conn[%v] [err=%v]", c.RemoteAddr().String(), err.Error())
			}
			return gnet.Close
		}
		s.log.Infof("conn[%v] receive [op=%v] [msg=%v]", c.RemoteAddr().String(), op, string(msg))
		// 使用handle将conn传到上层
		err = s.onMessageHandler(c, msg)
		if err != nil {
			return gnet.Close
		}
	}
	return gnet.None
}

func (s *Server) OnTick() (delay time.Duration, action gnet.Action) {
	s.log.Infof("[connected-count=%v]", atomic.LoadInt64(&s.connected))
	// TODO 定时处理
	return 10 * time.Second, gnet.None
}
