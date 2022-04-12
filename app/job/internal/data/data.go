package data

import (
	"arod-im/app/job/internal/conf"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/segmentio/kafka-go"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewConsumer, NewJobRepo)

// Data .
type Data struct {
	consumer *kafka.Reader
}

// NewData
func NewData(c *conf.Data, consumer *kafka.Reader, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{consumer: consumer}, cleanup, nil
}

func NewConsumer(c *conf.Data) *kafka.Reader {
	config := kafka.ReaderConfig{
		Brokers:   []string{"101.43.63.229:9092"},
		Topic:     "arod-im-push-topic",
		Partition: 0,
	}
	kafka.NewReader(config)
	//consumer, err := cluster.NewConsumer(c.Kafka.Brokers, c.Kafka.Group, []string{c.Kafka.Topic}, config)
	consumer := kafka.NewReader(config)
	return consumer
}
