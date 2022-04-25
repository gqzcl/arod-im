package service

import (
	v1 "arod-im/api/logic/v1"
	"arod-im/pkg/transport/websocket"
	"context"
	"encoding/json"
	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"github.com/panjf2000/gnet/v2"
)

// SetWebsocketServer 初始化ws server
func (s *ConnectorService) SetWebsocketServer(ws *websocket.Server) {
	s.ws = ws
}

type WebsocketProto struct {
	Uid     string `json:"uid"`
	GroupId string `json:"group_id"`
	RoomId  string `json:"room_id"`
}

// OnMessageHandle  接收ws对端的消息，由websocket server调用
func (s ConnectorService) OnMessageHandler(c gnet.Conn, message []byte) error {
	address := c.RemoteAddr().String()
	s.log.Infof("[%s] Payload: %s\n", c.RemoteAddr().String(), string(message))

	wsutil.WriteServerMessage(c, ws.OpText, []byte("receive success!"))

	var cookie WebsocketProto
	if err := json.Unmarshal(message, &cookie); err != nil {
		s.log.Error("Error unmarshalling proto json %v", err)
		return nil
	}
	// TODO 处理心跳
	// note： 初次连接，鉴权，将地址存入redis
	connect, err := s.LogicClient[0].Connect(context.Background(), &v1.ConnectReq{
		Server:  s.Address,
		Uid:     cookie.Uid,
		Address: address,
		Token:   nil,
	})
	if err != nil {
		return err
	}

	if !connect.Success {
		// TODO 连接操作失败
	} else {
		// 存储连接
		s.bc.AddCh(address, c)
	}

	bufProto, _ := json.Marshal(&cookie)
	var msg websocket.Message
	msg.Body = bufProto
	s.log.Info("msgBody is ", cookie.Uid, "+", cookie.GroupId)

	return nil
}

func (s *ConnectorService) OnCloseHandler(c gnet.Conn) {
	_ = c.RemoteAddr().String()
	//s.log.Info("id", address, register)
	//if register {
	//} else {
	//	s.log.Infof("Connection with ID %s was removed", address)
	//	s.bc.DeleteCh(address)
	//}
}
