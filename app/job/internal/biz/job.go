package biz

import (
	"context"

	v1 "arod-im/api/logic/v1"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

var (
	// ErrUserNotFound is user not found.
	ErrUserNotFound = errors.NotFound(v1.ErrorReason_USER_NOT_FOUND.String(), "user not found")
)

// JobRepo  is a Job repo.
type JobRepo interface {
	Consumer(ctx context.Context) error
	SingleSend(ctx context.Context, msg []byte)
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
func (uc *JobUsecase) PushMsg(ctx context.Context, msg []byte) (err error) {
	return nil
}
