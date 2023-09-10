// Copyright 2022 gqzcl <gqzcl@qq.com>. All rights reserved.
// Use of this source code is governed by a MIT style

package service

import (
	v1 "arod-im/api/logic/v1"
	"context"
	"time"
)

// SingleSend send a single message
func (s *MessageService) SingleSend(ctx context.Context, request *v1.SingleSendRequest) (*v1.SendReplay, error) {

	s.log.WithContext(ctx).Debug("SingleSend receive a msg from user:", request.Uid, "content:", request.MsgBody)

	seq, err := s.sc.PushMsg(ctx, request.Uid, request.Cid, request.MsgBody)
	if err != nil {
		return &v1.SendReplay{
			ActionStatus: "FAIL",
			ErrorInfo:    err.Error(),
			ErrorCode:    90001,
		}, err
	}

	return &v1.SendReplay{
		ActionStatus: "OK",
		MsgTime:      time.Now().UnixMicro(),
		MsgSeq:       seq,
	}, nil
}

func (s *MessageService) SingleRecall(ctx context.Context, request *v1.SingleRecallRequest) (*v1.RecallReplay, error) {
	return &v1.RecallReplay{
		ActionStatus: "OK",
		ErrorInfo:    "success",
		ErrorCode:    0,
	}, nil
}
