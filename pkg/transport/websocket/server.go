package websocket

import (
	"context"
	"crypto/tls"
	"net"
	"net/url"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/gobwas/ws"
	"github.com/panjf2000/gnet/v2"
)

type Message struct {
	Body []byte
}

var (
	_ transport.Server     = (*Server)(nil)
	_ transport.Endpointer = (*Server)(nil)
)

type wsCodec struct {
	ws bool
}

type Server struct {
	gnet.BuiltinEventEngine
	lis         net.Listener
	tlsConf     *tls.Config
	endpoint    *url.URL
	strictSlash bool

	eng       gnet.Engine
	err       error
	network   string
	address   string
	path      string
	connected int64
	timeout   time.Duration

	log *log.Helper

	onMessageHandler OnMessageHandler
	onCloseHandler   OnCloseHandler

	upgrader *ws.Upgrader
}

func NewServer(opts ...ServerOption) *Server {
	srv := &Server{
		network:     "tcp",
		address:     "7700",
		timeout:     2 * time.Second,
		strictSlash: true,
		log:         log.NewHelper(log.With(log.GetLogger(), "module", "ws-server")),

		upgrader: &ws.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		},
	}

	srv.init(opts...)
	srv.endpoint, srv.err = url.Parse(srv.address)
	return srv
}

func (s *Server) Name() string {
	return "websocket"
}

// 运行options
func (s *Server) init(opts ...ServerOption) {
	for _, o := range opts {
		o(s)
	}
}

// Endpoint 实现transport.Endpointer接口
func (s *Server) Endpoint() (*url.URL, error) {
	if s.err != nil {
		return nil, s.err
	}
	return s.endpoint, nil
}

// Start 实现transport.Server接口
func (s *Server) Start(ctx context.Context) error {

	s.log.Infof("[websocket] server listening on: %s", s.address)

	err := gnet.Run(s, s.address,
		gnet.WithMulticore(true),
		gnet.WithReusePort(true),
		gnet.WithTicker(true))
	return err
}

// Stop 实现transport.Server接口
func (s *Server) Stop(ctx context.Context) error {
	s.log.Info("[websocket] server stopping")
	return gnet.Stop(ctx, s.address)
}
