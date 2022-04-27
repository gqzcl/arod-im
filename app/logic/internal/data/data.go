package data

import (
	"arod-im/app/logic/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/gomodule/redigo/redis"
	"github.com/google/wire"
	kafka "gopkg.in/Shopify/sarama.v1"
	"time"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewKafkaPub, NewRedis, NewSingleRepo, NewGroupRepo, NewRoomRepo, NewConnectRepo)

// Data .
type Data struct {
	// TODO wrapped database client
	kafkaPub    kafka.SyncProducer
	redis       *redis.Pool
	redisExpire time.Duration
}

// NewData .
func NewData(c *conf.Data, logger log.Logger, kafkaPub kafka.SyncProducer, redis *redis.Pool) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{kafkaPub: kafkaPub, redis: redis, redisExpire: c.Redis.Expire.AsDuration()}, cleanup, nil
}

func NewKafkaPub(c *conf.Data) kafka.SyncProducer {
	kc := kafka.NewConfig()
	kc.Producer.RequiredAcks = kafka.WaitForAll
	kc.Producer.Retry.Max = 10 // 重试次数10
	kc.Producer.Return.Successes = true
	pub, err := kafka.NewSyncProducer([]string{"101.43.63.229:9092"}, kc)
	if err != nil {
		panic(err)
	}
	return pub
}

func NewRedis(c *conf.Data) *redis.Pool {
	r := &redis.Pool{
		MaxIdle:     int(c.Redis.Idle),
		MaxActive:   int(c.Redis.Active),
		IdleTimeout: c.Redis.IdleTimeout.AsDuration(),
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial(c.Redis.Network, c.Redis.Addr,
				redis.DialConnectTimeout(c.Redis.DailTimeout.AsDuration()),
				redis.DialReadTimeout(c.Redis.ReadTimeout.AsDuration()),
				redis.DialWriteTimeout(c.Redis.WriteTimeout.AsDuration()),
				redis.DialPassword(c.Redis.Auth),
			)
			if err != nil {
				panic(err)
			}
			return conn, nil
		},
	}
	return r
}
