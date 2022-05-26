// Copyright 2022 gqzcl <gqzcl@qq.com>. All rights reserved.
// Use of this source code is governed by a MIT style

package kafka

import (
	"crypto/tls"
	"github.com/go-kratos/kratos/v2/log"
)

type ServerOption func(o *Server)

func NewReader(brokers []string, topic string, opts ...ConfigOption) ServerOption {
	return func(s *Server) {
		s.readerConfig.Brokers = brokers
		s.readerConfig.Topic = topic
		for _, option := range opts {
			option(s.readerConfig)
		}
	}
}

func OnMessage(handler Handler) ServerOption {
	return func(s *Server) {
		s.handler = handler
	}
}

func Logger(logger log.Logger) ServerOption {
	return func(s *Server) {
		s.log = log.NewHelper(logger)
	}
}

func TLSConfig(c *tls.Config) ServerOption {
	return func(s *Server) {
	}
}
