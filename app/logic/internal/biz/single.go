// Copyright 2022 gqzcl <gqzcl@qq.com>. All rights reserved.
// Use of this source code is governed by a MIT style

package biz

import (
	jobV1 "arod-im/api/job/v1"
	v1 "arod-im/api/logic/v1"
	"arod-im/pkg/rambler"
	"context"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

var (
	// ErrUserNotFound is user not found.
	ErrUserNotFound = errors.NotFound(v1.ErrorReason_USER_NOT_FOUND.String(), "user not found")
)

type Greeter struct {
}

//// MessageBody  is a Message model.
//type MessageBody struct {
//	Address map[string]string `json:"address"`
//	Body    []*v1.MsgBody     `json:"body"`
//}

// SingleRepo is a Single repo.
type SingleRepo interface {
	Push(ctx context.Context, sessionId string, msg *jobV1.SingleSendMsg) (err error)
	GetUserAddress(ctx context.Context, uid string) (map[string]string, error)
}

// SingleUsecase is a Single usecase.
type SingleUsecase struct {
	single  SingleRepo
	log     *log.Helper
	rambler *rambler.Rambler
}

// NewSingleUsecase new a Single usecase.
func NewSingleUsecase(single SingleRepo, logger log.Logger) *SingleUsecase {
	return &SingleUsecase{single: single, log: log.NewHelper(logger), rambler: rambler.NewRambler()}
}

// PushMsg push a msg to data.
func (sc *SingleUsecase) PushMsg(ctx context.Context, uid, cid string, msg []*jobV1.MsgBody) (string, error) {
	seq := sc.rambler.GetSeqID([]byte(uid + cid))
	addrs, err := sc.single.GetUserAddress(ctx, cid)
	if err != nil {
		sc.log.WithContext(ctx).Error(err)
	}
	sc.log.WithContext(ctx).Info(addrs)
	message := &jobV1.SingleSendMsg{
		Server: addrs,
		Seq:    seq,
		Msg:    msg,
	}
	if err := sc.single.Push(ctx, uid+cid, message); err != nil {
		sc.log.WithContext(ctx).Info("Error in PushMsg() : ", err)
		return "", err
	}
	sc.log.WithContext(ctx).Info("PushMsg Success: ", seq)
	return seq, nil
}

func (sc *SingleUsecase) RecallMsg(ctx context.Context, key string) error {
	return nil
}
