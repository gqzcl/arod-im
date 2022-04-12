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

// Greeter is a Greeter model.
type Greeter struct {
	Hello string
}

// GreeterRepo is a Greater repo.
type JobRepo interface {
	Consumer(ctx context.Context) error
}

// GreeterUsecase is a Greeter usecase.
type JobUsecase struct {
	job JobRepo
	log *log.Helper
}

// NewGreeterUsecase new a Greeter usecase.
func NewJobUsecase(job JobRepo, logger log.Logger) *JobUsecase {
	return &JobUsecase{job: job, log: log.NewHelper(logger)}
}

// CreateGreeter creates a Greeter, and returns the new Greeter.
func (uc *JobUsecase) Consumer(ctx context.Context) (err error) {
	for {
		err := uc.job.Consumer(ctx)
		if err != nil {
			return err
		}
	}
}
