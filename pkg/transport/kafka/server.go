package kafka

import (
	"arod-im/pkg/broker"
	"arod-im/pkg/broker/kafka"
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport"

	"net/url"
	"strings"
	"sync"
)

var (
	_ transport.Server     = (*Server)(nil)
	_ transport.Endpointer = (*Server)(nil)
)

type SubscriberMap map[string]broker.Subscriber

type SubscribeOption struct {
	handler broker.Handler
	opts    []broker.SubscribeOption
}
type SubscribeOptionMap map[string]*SubscribeOption

type Server struct {
	broker.Broker
	bOpts []broker.Option

	subscribers    SubscriberMap
	subscriberOpts SubscribeOptionMap

	sync.RWMutex
	started bool

	log     *log.Helper
	baseCtx context.Context
	err     error
}

func NewServer(opts ...ServerOption) *Server {
	srv := &Server{
		baseCtx:        context.Background(),
		log:            log.NewHelper(log.GetLogger()),
		subscribers:    SubscriberMap{},
		subscriberOpts: SubscribeOptionMap{},
		bOpts:          []broker.Option{},
		started:        false,
	}

	srv.init(opts...)
	srv.Broker = kafka.NewBroker(srv.bOpts...)

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

func (s *Server) Endpoint() (*url.URL, error) {
	if s.err != nil {
		return nil, s.err
	}

	addr := s.Address()
	if !strings.HasPrefix(addr, "tcp://") {
		addr = "tcp://" + addr
	}

	return url.Parse(addr)
}

func (s *Server) Start(ctx context.Context) error {
	if s.err != nil {
		return s.err
	}

	if s.started {
		return nil
	}

	s.err = s.Connect()
	if s.err != nil {
		return s.err
	}

	s.log.Infof("[kafka] server listening on: %s", s.Address())

	s.err = s.doRegisterSubscriberMap()
	if s.err != nil {
		return s.err
	}

	s.baseCtx = ctx
	s.started = true

	return nil
}

func (s *Server) Stop(_ context.Context) error {
	if s.started == false {
		return nil
	}
	s.log.Info("[kafka] server stopping")

	for _, v := range s.subscribers {
		_ = v.Unsubscribe()
	}
	s.subscribers = SubscriberMap{}
	s.subscriberOpts = SubscribeOptionMap{}

	s.started = false
	return s.Disconnect()
}

func (s *Server) RegisterSubscriber(topic string, h broker.Handler, opts ...broker.SubscribeOption) error {
	s.Lock()
	defer s.Unlock()

	if s.started {
		return s.doRegisterSubscriber(topic, h, opts...)
	} else {
		s.subscriberOpts[topic] = &SubscribeOption{handler: h, opts: opts}
	}
	return nil
}

func (s *Server) doRegisterSubscriber(topic string, h broker.Handler, opts ...broker.SubscribeOption) error {
	sub, err := s.Subscribe(topic, h, opts...)
	if err != nil {
		return err
	}

	s.subscribers[topic] = sub

	return nil
}

func (s *Server) doRegisterSubscriberMap() error {
	for topic, opt := range s.subscriberOpts {
		_ = s.doRegisterSubscriber(topic, opt.handler, opt.opts...)
	}
	s.subscriberOpts = SubscribeOptionMap{}
	return nil
}
