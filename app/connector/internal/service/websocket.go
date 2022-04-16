package service

import (
	"arod-im/pkg/transport/websocket"
	"encoding/json"

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

	var proto WebsocketProto
	if err := json.Unmarshal(message, &proto); err != nil {
		s.log.Error("Error unmarshalling proto json %v", err)
		return nil
	}
	s.bc.AddCh(address, c)
	// TODO 将地址存入redis
	bufProto, _ := json.Marshal(&proto)
	var msg websocket.Message
	msg.Body = bufProto
	s.log.Info("msgBody is ", proto.Uid, "+", proto.GroupId)
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
