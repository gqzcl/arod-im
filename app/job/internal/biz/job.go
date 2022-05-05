package biz

import (
	"context"

	jobV1 "arod-im/api/job/v1"
	"github.com/go-kratos/kratos/v2/log"
)

var (
// ErrUserNotFound is user not found.
// ErrUserNotFound = errors.NotFound(v1.ErrorReason_USER_NOT_FOUND.String(), "user not found")
)

// JobRepo  is a Job repo.
type JobRepo interface {
	SingleSend(ctx context.Context, server, address, seq string, msg []*jobV1.MsgBody)
}

// JobUsecase is a Job usecase.
type JobUsecase struct {
	job JobRepo
	log *log.Helper
}

// NewJobUsecase  new a Job usecase.
func NewJobUsecase(job JobRepo, logger log.Logger) *JobUsecase {
	return &JobUsecase{job: job, log: log.NewHelper(logger)}
}

// PushMsg push msg to connector
func (uc *JobUsecase) PushMsg(ctx context.Context, server, address, seq string, msg []*jobV1.MsgBody) (err error) {
	uc.log.WithContext(ctx).Debugf("Push Msg to %s->%s ,the content is %s", server, address, msg)
	uc.job.SingleSend(ctx, server, address, seq, msg)
	return nil
}
