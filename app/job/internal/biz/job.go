// Copyright 2022 gqzcl <gqzcl@qq.com>. All rights reserved.
// Use of this source code is governed by a MIT style

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
	SingleSend(ctx context.Context, address, server, senderId, seq string, msg []*jobV1.MsgBody) error
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
func (uc *JobUsecase) PushMsg(ctx context.Context, msg *jobV1.SingleSendMsg) (err error) {
	uc.log.WithContext(ctx).Debugf("Push Msg the content is %s", msg)
	for address := range msg.Server {
		err := uc.job.SingleSend(ctx, address, msg.Server[address], msg.SenderId, msg.Seq, msg.Msg)
		if err != nil {
			return err
		}
	}
	return nil
}
