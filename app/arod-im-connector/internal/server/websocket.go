// Copyright 2022 gqzcl <gqzcl@qq.com>. All rights reserved.
// Use of this source code is governed by a MIT style

package server

import (
	"arod-im/app/arod-im-connector/internal/conf"
	"arod-im/app/arod-im-connector/internal/service"
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

	//s.SetWebsocketServer(srv)

	return srv
}
