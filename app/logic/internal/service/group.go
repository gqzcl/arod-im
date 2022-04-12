package service

import (
	v1 "arod-im/api/logic/v1"
	"context"
)

func (s *MessageService) GroupSend(ctx context.Context, request *v1.GroupSendRequest) (*v1.SendReplay, error) {
	return &v1.SendReplay{
		ActionStatus: "OK",
		ErrorInfo:    "success",
		ErrorCode:    0,
		MsgTime:      132165465,
		MsgSeq:       "sadasda",
	}, nil
}

func (s *MessageService) GroupSendMention(ctx context.Context, request *v1.GroupSendMentionRequest) (*v1.SendReplay, error) {
	return &v1.SendReplay{
		ActionStatus: "OK",
		ErrorInfo:    "success",
		ErrorCode:    0,
		MsgTime:      132165465,
		MsgSeq:       "sadasda",
	}, nil
}

func (s *MessageService) GroupRecall(ctx context.Context, request *v1.GroupRecallRequest) (*v1.RecallReplay, error) {
	return &v1.RecallReplay{
		ActionStatus: "OK",
		ErrorInfo:    "success",
		ErrorCode:    0,
	}, nil
}
