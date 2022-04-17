package server

import (
	"arod-im/app/job/internal/conf"
	"arod-im/app/job/internal/service"
	"arod-im/pkg/transport/kafka"
)

func NewKafkaServer(c *conf.Server, s *service.JobService) *kafka.Server {
	//ctx := context.Background()
	srv := kafka.NewServer(
		kafka.NewReader(c.Kafka.Brokers, c.Kafka.Topic),
		kafka.OnMessage(s.OnMessage),
	)
	return srv
}
