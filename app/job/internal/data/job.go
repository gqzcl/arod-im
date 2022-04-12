package data

import (
	"arod-im/app/job/internal/biz"
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type MsgBody struct {
}

type jobRepo struct {
	data *Data
	log  *log.Helper
}

// NewJobRepo NewGreeterRepo .
func NewJobRepo(data *Data, logger log.Logger) biz.JobRepo {
	return &jobRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (j *jobRepo) Consumer(ctx context.Context) error {
	m, err := j.data.consumer.ReadMessage(ctx)
	if err != nil {
		return err
	}
	j.log.WithContext(ctx).Infof("message at offset %d: %s = %s\n", m.Offset, string(m.Key), string(m.Value))
	return nil
}
