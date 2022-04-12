package service

import (
	v1 "arod-im/api/job/v1"
	"arod-im/app/job/internal/biz"
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

// GreeterService is a greeter service.
type JobService struct {
	v1.UnimplementedJobServer

	jc  *biz.JobUsecase
	log *log.Helper
}

// NewGreeterService new a greeter service.
func NewJobService(jc *biz.JobUsecase, logger log.Logger) *JobService {
	return &JobService{
		jc:  jc,
		log: log.NewHelper(logger),
	}
}

func NewConsumerP1(js *JobService) error {
	js.GetMsg(context.Background())
	return nil
}

// SayHello implements helloworld.GreeterServer.
func (j *JobService) GetMsg(ctx context.Context) {
	err := j.jc.Consumer(ctx)
	if err != nil {
		return
	}
}
