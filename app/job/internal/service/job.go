package service

import (
	v1 "arod-im/api/job/v1"
	logicV1 "arod-im/api/logic/v1"
	"arod-im/app/job/internal/biz"
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/tx7do/kratos-transport/broker"
	"net/http"
)

// GreeterService is a greeter service.
type JobService struct {
	*http.Server
	v1.UnimplementedJobServer

	kb  broker.Broker
	jc  *biz.JobUsecase
	log *log.Helper
}

// NewGreeterService new a greeter service.
func NewJobService(jc *biz.JobUsecase, logger log.Logger) *JobService {
	j := &JobService{

		jc:  jc,
		log: log.NewHelper(logger),
	}
	go j.GetMsg(context.Background())
	return j
}

func (j *JobService) SetKafkaBroker(b broker.Broker) {
	j.kb = b
}

func (j *JobService) NewConsumer(event broker.Event) error {
	fmt.Println("NewConsumer() Topic: ", event.Topic(), " Payload: ", string(event.Message().Body))
	var m []*logicV1.SingleSendRequest_MsgBody
	err := json.Unmarshal(event.Message().Body, &m)
	j.log.Info(err)
	//j.GetService()
	if err != nil {
		return err
	}
	j.log.Info(m[0].MsgType, m[0].MsgContent)
	return nil
}

// SayHello implements helloworld.GreeterServer.
func (j *JobService) GetMsg(ctx context.Context) {
	err := j.jc.Consumer(ctx)
	if err != nil {
		j.log.Error("kafka 消费失败", err)
		return
	}
}
