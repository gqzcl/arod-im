package service

import (
	v1 "arod-im/api/logic/v1"
	"context"
	"time"
)

// SingleSend send a single message
func (s *MessageService) SingleSend(ctx context.Context, request *v1.SingleSendRequest) (*v1.SendReplay, error) {
	s.log.WithContext(ctx).Info("Single Send a msg from user", request.Uid, "content:", request.MsgBody)
	seq, err := s.sc.PushMsg(ctx, request.Uid, request.Cid, request.MsgBody)
	if err != nil {
		return &v1.SendReplay{
			ActionStatus: "FAIL",
			ErrorInfo:    err.Error(),
			ErrorCode:    90001,
		}, err
	}
	// TODO 修改一下时间
	return &v1.SendReplay{
		ActionStatus: "OK",
		ErrorInfo:    "",
		ErrorCode:    0,
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
