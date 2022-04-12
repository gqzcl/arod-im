package biz

import (
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

// MessageBody  is a Message model.
type MessageBody struct {
	msgType    string
	msgContent []byte
}

// SingleRepo is a Single repo.
type SingleRepo interface {
	Push(ctx context.Context, sessionId string, msg []*v1.SingleSendRequest_MsgBody) (err error)
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
func (sc *SingleUsecase) PushMsg(ctx context.Context, uid, cid string, msg []*v1.SingleSendRequest_MsgBody) (string, error) {
	seq := sc.rambler.GetSeqID([]byte(uid + cid))

	//m := make([]*MessageBody, 0)
	//for i := range msg {
	//	b, err := msg[i].GetMsgContent().MarshalJSON()
	//	if err != nil {
	//		sc.log.WithContext(ctx).Infof("err: %v", err)
	//		return "", err
	//	}
	//	m = append(m, &MessageBody{msgType: msg[i].MsgType.String(), msgContent: b})
	//}
	if err := sc.single.Push(ctx, uid+cid, msg); err != nil {
		sc.log.WithContext(ctx).Info("Error in PushMsg() : ", err)
		return "", err
	}
	sc.log.WithContext(ctx).Info("PushMsg Success: ", seq)
	return seq, nil
}

func (sc *SingleUsecase) RecallMsg(ctx context.Context, key string) error {
	return nil
}
