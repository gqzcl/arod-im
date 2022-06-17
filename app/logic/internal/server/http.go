// Copyright 2022 gqzcl <gqzcl@qq.com>. All rights reserved.
// Use of this source code is governed by a MIT style

package server

import (
	v1 "arod-im/api/logic/v1"
	"arod-im/app/logic/internal/conf"
	"arod-im/app/logic/internal/service"
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/transport/http"
)

func SkipRouterMatch() selector.MatchFunc {
	whiteList := make(map[string]struct{})
	whiteList["/api.logic.v1.Logic/Login"] = struct{}{}
	return func(ctx context.Context, operation string) bool {
		if _, ok := whiteList[operation]; ok {
			return false
		}
		return true
	}
}

// NewHTTPServer new a HTTP server.
func NewHTTPServer(c *conf.Server, greeter *service.MessageService, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			selector.Server(JWTAuth()).Match(SkipRouterMatch()).Build(),
			// jwt.Server(
			// 	func(token *jwtv4.Token) (interface{}, error) {
			// 		return []byte(c.SecretKey), nil
			// 	},
			// 	jwt.WithClaims(func() jwtv4.Claims {
			// 		return &jwtv4.StandardClaims{}
			// 	}),
			// 	jwt.WithSigningMethod(jwtv4.SigningMethodES256),
			// ),
		),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	v1.RegisterLogicHTTPServer(srv, greeter)
	return srv
}
