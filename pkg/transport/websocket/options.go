package websocket

import (
	"crypto/tls"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"net"
	"time"
)

type ServerOption func(o *Server)

func Address(addr string) ServerOption {
	return func(s *Server) {
		s.address = fmt.Sprintf("tcp://127.0.0.1:%s", addr)
	}
}

func Path(path string) ServerOption {
	return func(s *Server) {
		s.path = path
	}
}

func Timeout(timeout time.Duration) ServerOption {
	return func(s *Server) {
		s.timeout = timeout
	}
}

func OnMessageHandle(h OnMessageHandler) ServerOption {
	return func(s *Server) {
		s.onMessageHandler = h
	}
}

func OnCloseHandle(h OnCloseHandler) ServerOption {
	return func(s *Server) {
		s.onCloseHandler = h
	}
}

func Logger(logger log.Logger) ServerOption {
	return func(s *Server) {
		s.log = log.NewHelper(logger)
	}
}

func TLSConfig(c *tls.Config) ServerOption {
	return func(o *Server) {
		o.tlsConf = c
	}
}

func Listener(lis net.Listener) ServerOption {
	return func(s *Server) {
		s.lis = lis
	}
}
