// Copyright 2022 gqzcl <gqzcl@qq.com>. All rights reserved.
// Use of this source code is governed by a MIT style

package server

// NewHTTPServer new a HTTP server.
//func NewHTTPServer(c *conf.Server, greeter *service.GreeterService, logger log.Logger) *http.Server {
//	var opts = []http.ServerOption{
//		http.Middleware(
//			recovery.Recovery(),
//		),
//	}
//	if c.Http.Network != "" {
//		opts = append(opts, http.Network(c.Http.Network))
//	}
//	if c.Http.Addr != "" {
//		opts = append(opts, http.Address(c.Http.Addr))
//	}
//	if c.Http.Timeout != nil {
//		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
//	}
//	srv := http.NewServer(opts...)
//	v1.RegisterConnectorServer(srv, greeter)
//	return srv
//}
