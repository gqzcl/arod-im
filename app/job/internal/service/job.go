package service

import (
	v1 "arod-im/api/job/v1"
	logicV1 "arod-im/api/logic/v1"
	"arod-im/app/job/internal/biz"
	"context"
	"encoding/json"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/segmentio/kafka-go"
)

// JobService  is a Job service.
type JobService struct {
	v1.UnimplementedJobServer
	jc  *biz.JobUsecase
	log *log.Helper
}

// NewJobService  new a Job service.
func NewJobService(jc *biz.JobUsecase, logger log.Logger) *JobService {
	j := &JobService{
		jc:  jc,
		log: log.NewHelper(logger),
	}
	return j
}

// OnMessage message format: serverId , address , senderId , msg[ id , content]
func (j *JobService) OnMessage(ctx context.Context, message kafka.Message) error {
	j.log.WithContext(ctx).Debugf("message at topic/partition/offset %v/%v/%v: %s = %s\n",
		message.Topic, message.Partition, message.Offset, string(message.Key), string(message.Value))

	var m []*logicV1.SingleSendRequest_MsgBody
	err := json.Unmarshal(message.Value, &m)
	if err != nil {
		j.log.WithContext(ctx).Error(err)
	}

	err = j.jc.PushMsg(ctx, message.Value)
	if err != nil {
		j.log.WithContext(ctx).Error(err)
	}

	// TODO 回复ack
	// idea 同步方式，处理完一条消息再处理下一条消息，直接发送ack即可
	// 异步方式，还没处理完就处理下一条，可能出现处理失败的问题
	return nil
}
