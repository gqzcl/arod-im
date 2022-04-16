package server

import (
	"arod-im/app/job/internal/conf"
	"arod-im/app/job/internal/service"
	"context"
	"github.com/tx7do/kratos-transport/transport/kafka"
)

func NewKafkaServer(c *conf.Server, s *service.JobService) *kafka.Server {
	ctx := context.Background()
	srv := kafka.NewServer(
		kafka.Address(c.Kafka.Addrs[0]),
	)
	s.SetKafkaBroker(srv)
	s.GetMsg(ctx)
	return srv
}
