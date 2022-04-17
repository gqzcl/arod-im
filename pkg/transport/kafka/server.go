package kafka

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/segmentio/kafka-go"

	"net/url"
	"strings"
	"sync"
)

var (
	_ transport.Server     = (*Server)(nil)
	_ transport.Endpointer = (*Server)(nil)
)

type Handler func(ctx context.Context, message kafka.Message) error

type Server struct {
	handler Handler

	sync.RWMutex
	started bool

	reader       *kafka.Reader
	readerConfig kafka.ReaderConfig

	log     *log.Helper
	baseCtx context.Context
	err     error
}

func NewServer(opts ...ServerOption) *Server {
	srv := &Server{
		baseCtx: context.Background(),
		log:     log.NewHelper(log.GetLogger()),
		started: false,
	}

	srv.init(opts...)
	srv.reader = kafka.NewReader(srv.readerConfig)

	return srv
}

func (s *Server) init(opts ...ServerOption) {
	for _, o := range opts {
		o(s)
	}
}

func (s *Server) Name() string {
	return "kafka"
}

func (s *Server) Consumer(ctx context.Context) {
	for {
		m, err := s.reader.ReadMessage(ctx)
		err = s.handler(ctx, m)
		if err != nil {
			s.log.WithContext(ctx).Error(err)
		}
		// TODO 如果配置了不自动提交，则在成功操作后进行提交offset
	}
}

// Endpoint 实现transport.Endpointer接口
func (s *Server) Endpoint() (*url.URL, error) {
	if s.err != nil {
		return nil, s.err
	}

	addr := s.readerConfig.Brokers[0]
	if !strings.HasPrefix(addr, "tcp://") {
		addr = "tcp://" + addr
	}
	return url.Parse(addr)
}

// Start 实现transport.Server接口
func (s *Server) Start(ctx context.Context) error {
	// TODO
	if s.err != nil {
		return s.err
	}

	if s.started {
		return nil
	}

	s.log.Infof("[kafka] comsumer listening on: %s", s.readerConfig.Brokers)

	if s.err != nil {
		return s.err
	}

	s.baseCtx = ctx
	s.started = true
	go s.Consumer(ctx)
	return nil
}

// Stop 实现transport.Server接口
func (s *Server) Stop(_ context.Context) error {
	if s.started == false {
		return nil
	}
	s.log.Info("[kafka] server stopping")

	return s.reader.Close()
}
