package service

import (
	v1 "arod-im/api/job/v1"
	"arod-im/app/job/internal/biz"
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/segmentio/kafka-go"
	"net/http"
)

// GreeterService is a greeter service.
type JobService struct {
	*http.Server
	v1.UnimplementedJobServer
	jc  *biz.JobUsecase
	log *log.Helper
}

// NewGreeterService new a greeter service.
func NewJobService(jc *biz.JobUsecase, logger log.Logger) *JobService {
	j := &JobService{

		jc:  jc,
		log: log.NewHelper(logger),
	}
	return j
}

func (j *JobService) OnMessage(ctx context.Context, message kafka.Message) error {
	j.log.WithContext(ctx).Infof("message at topic/partition/offset %v/%v/%v: %s = %s\n", message.Topic, message.Partition, message.Offset, string(message.Key), string(message.Value))
	return nil
}

//func (j *JobService) NewConsumer(event broker.Event) error {
//	fmt.Println("NewConsumer() Topic: ", event.Topic(), " Payload: ", string(event.Message().Body))
//	var m []*logicV1.SingleSendRequest_MsgBody
//	err := json.Unmarshal(event.Message().Body, &m)
//	j.log.Info(err)
//	//j.GetService()
//	if err != nil {
//		return err
//	}
//	j.log.Info(m[0].MsgType, m[0].MsgContent)
//	return nil
//}

// SayHello implements helloworld.GreeterServer.
func (j *JobService) GetMsg(ctx context.Context) {
	err := j.jc.Consumer(ctx)
	if err != nil {
		j.log.Error("kafka 消费失败", err)
		return
	}
}
