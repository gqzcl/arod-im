package dao

import (
	"context"
	"time"

	"github.com/gomodule/redigo/redis"
	"github.com/gqzcl/gim/internal/logic/conf"
	kafka "gopkg.in/Shopify/sarama.v1"
)

type Dao struct {
	c           *conf.Config
	kafkaPub    kafka.SyncProducer
	redis       *redis.Pool
	redisExpire int32
}

func New(c *conf.Config) *Dao {
	d := &Dao{
		c:           c,
		kafkaPub:    newKafkaPub(c.Kafka),
		redis:       newRedis(c.Redis),
		redisExpire: int32(c.Redis.Expire) / int32(time.Second),
	}
	return d
}

func newKafkaPub(c *conf.Kafka) kafka.SyncProducer {
	kc := kafka.NewConfig()
	// wait for all in-sync replicas to ack the message
	kc.Producer.RequiredAcks = kafka.WaitForAll
	// retry up to 10 times to produce the message
	kc.Producer.Retry.Max = 10
	// return the successfully delivered messages
	kc.Producer.Return.Successes = true
	pud, err := kafka.NewSyncProducer(c.Brokers, kc)
	if err != nil {
		panic(err)
	}
	return pud
}

func newRedis(c *conf.Redis) *redis.Pool {
	return &redis.Pool{
		// 池中空闲连接的最大数量
		MaxIdle: c.Idle,
		// 连接池在给定时间分配的最大连接数。当为零时，池中的连接数没有限制。
		MaxActive:   c.Active,
		IdleTimeout: c.IdleTimeout,
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial(c.Network, c.Addr,
				redis.DialConnectTimeout(c.DialTimeout),

				redis.DialReadTimeout(c.ReadTimeout),

				redis.DialWriteTimeout(c.WriteTimeout),

				redis.DialPassword(c.Auth),
			)
			if err != nil {
				return nil, err
			}
			return conn, nil
		},
	}
}

// Close close the redis resource
func (d *Dao) Close() error {
	return d.redis.Close()
}

// Ping dao ping
func (d *Dao) Ping(c context.Context) error {
	return d.pingRedis(c)
}

func (d *Dao) pingRedis(c context.Context) error {
	conn := d.redis.Get()
	_, err := conn.Do("SET", "PING", "PONG")
	conn.Close()
	return err
}
