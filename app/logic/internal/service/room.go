package service

import (
	v1 "arod-im/api/logic/v1"
	"context"
)

func (s *MessageService) RoomSend(ctx context.Context, request *v1.GroupSendRequest) (*v1.SendReplay, error) {
	return &v1.SendReplay{
		ActionStatus: "OK",
		ErrorInfo:    "success",
		ErrorCode:    0,
		MsgTime:      132165465,
		MsgSeq:       "sadasda",
	}, nil
}

func (s *MessageService) RoomBroadcast(ctx context.Context, request *v1.GroupSendRequest) (*v1.SendReplay, error) {
	return &v1.SendReplay{
		ActionStatus: "OK",
		ErrorInfo:    "success",
		ErrorCode:    0,
		MsgTime:      132165465,
		MsgSeq:       "sadasda",
	}, nil
}
