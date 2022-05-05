package data

import (
	v1 "arod-im/api/logic/v1"
	"arod-im/app/logic/internal/biz"
	"context"
	"github.com/golang/protobuf/proto"
	"gopkg.in/Shopify/sarama.v1"

	"github.com/go-kratos/kratos/v2/log"
)

type singleRepo struct {
	data *Data
	log  *log.Helper
}

// NewSingleRepo new a single repo
func NewSingleRepo(data *Data, logger log.Logger) biz.SingleRepo {
	return &singleRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *singleRepo) Push(ctx context.Context, sessionId string, msg *v1.SendMsg) (err error) {
	pushMsg := msg
	p, err := proto.Marshal(pushMsg)
	r.log.WithContext(ctx).Debugf("msg in Push :%s", p)
	m := &sarama.ProducerMessage{
		Key:   sarama.StringEncoder(sessionId),
		Topic: "arod-im-push-topic",
		Value: sarama.ByteEncoder(p),
	}
	if _, _, err = r.data.kafkaPub.SendMessage(m); err != nil {
		r.log.WithContext(ctx).Info(err)
		return err
	}
	return err
}

func (r *singleRepo) GetUserAddress(ctx context.Context, uid string) (addrs map[string]string, err error) {
	addrs, err = r.data.GetUserAddress(ctx, uid)
	// TODO deal err
	return
}
