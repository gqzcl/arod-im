// Copyright 2022 gqzcl <gqzcl@qq.com>. All rights reserved.
// Use of this source code is governed by a MIT style

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
//func (s *ConnectorService) SetWebsocketServer(ws *websocket.Server) {
//	s.ws = ws
//}

type WebsocketProto struct {
	Uid     string `json:"uid"`
	GroupId string `json:"group_id"`
	RoomId  string `json:"room_id"`
}

// TODO 处理只建立了连接没有鉴权的情况（清理掉）
// idea 定时五秒没有接收到鉴权消息则关闭连接
// OnMessageHandle  接收ws对端的消息，由websocket server调用
func (s *ConnectorService) OnMessageHandler(c gnet.Conn, message []byte) error {
	address := c.RemoteAddr().String()
	s.log.Infof("[%s] Payload: %s\n", c.RemoteAddr().String(), string(message))

	wsutil.WriteServerMessage(c, ws.OpText, []byte("receive success!"))

	var cookie WebsocketProto
	if err := json.Unmarshal(message, &cookie); err != nil {
		s.log.Error("Error unmarshalling proto json %v", err)
		return nil
	}
	// TODO 处理心跳
	// TODO 添加一个错误原因，如鉴权失败则断开连接
	// note： 初次连接，鉴权，将地址存入redis
	// c.Close(nil)
	success, err := s.StoreConnect(cookie.Uid, address)
	// 存储uid，关联uid与ws连接
	c.Context().(*websocket.WsContext).Uid = cookie.Uid
	s.log.Infof("成功连接")
	if err != nil {
		return err
	}

	if !success {
		// TODO 连接操作失败
		c.Close(nil)
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
	address := c.RemoteAddr().String()
	uid := c.Context().(*websocket.WsContext).Uid
	connect, _ := s.LogicClient.Disconnect(context.Background(), &v1.DisConnectReq{
		Server: s.Address,
		// TODO how 获得uid
		Uid:     uid,
		Address: address,
	})
	s.log.Infof("成功关闭连接")

	if !connect.Success {
		// TODO 关闭连接操作失败
	} else {
		// 存储连接
		s.bc.DeleteCh(address)
	}
}
