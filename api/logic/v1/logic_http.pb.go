// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// protoc-gen-go-http v2.3.1

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationLogicGetService = "/api.logic.v1.Logic/GetService"
const OperationLogicGroupRecall = "/api.logic.v1.Logic/GroupRecall"
const OperationLogicGroupSend = "/api.logic.v1.Logic/GroupSend"
const OperationLogicGroupSendMention = "/api.logic.v1.Logic/GroupSendMention"
const OperationLogicLogin = "/api.logic.v1.Logic/Login"
const OperationLogicRoomBroadcast = "/api.logic.v1.Logic/RoomBroadcast"
const OperationLogicRoomSend = "/api.logic.v1.Logic/RoomSend"
const OperationLogicSingleRecall = "/api.logic.v1.Logic/SingleRecall"
const OperationLogicSingleSend = "/api.logic.v1.Logic/SingleSend"

type LogicHTTPServer interface {
	GetService(context.Context, *GetServiceReq) (*GetServiceReplay, error)
	GroupRecall(context.Context, *GroupRecallRequest) (*RecallReplay, error)
	GroupSend(context.Context, *GroupSendRequest) (*SendReplay, error)
	GroupSendMention(context.Context, *GroupSendMentionRequest) (*SendReplay, error)
	Login(context.Context, *LoginReq) (*LoginReplay, error)
	RoomBroadcast(context.Context, *GroupSendRequest) (*SendReplay, error)
	RoomSend(context.Context, *GroupSendRequest) (*SendReplay, error)
	SingleRecall(context.Context, *SingleRecallRequest) (*RecallReplay, error)
	SingleSend(context.Context, *SingleSendRequest) (*SendReplay, error)
}

func RegisterLogicHTTPServer(s *http.Server, srv LogicHTTPServer) {
	r := s.Route("/")
	r.POST("v1/single/send", _Logic_SingleSend0_HTTP_Handler(srv))
	r.POST("v1/single/recall", _Logic_SingleRecall0_HTTP_Handler(srv))
	r.POST("v1/group/send", _Logic_GroupSend0_HTTP_Handler(srv))
	r.POST("v1/group/send_mention", _Logic_GroupSendMention0_HTTP_Handler(srv))
	r.POST("v1/group/recall", _Logic_GroupRecall0_HTTP_Handler(srv))
	r.POST("v1/room/send", _Logic_RoomSend0_HTTP_Handler(srv))
	r.POST("v1/room/broadcast", _Logic_RoomBroadcast0_HTTP_Handler(srv))
	r.POST("v1/login", _Logic_Login0_HTTP_Handler(srv))
	r.GET("v1/service", _Logic_GetService0_HTTP_Handler(srv))
}

func _Logic_SingleSend0_HTTP_Handler(srv LogicHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in SingleSendRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationLogicSingleSend)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.SingleSend(ctx, req.(*SingleSendRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*SendReplay)
		return ctx.Result(200, reply)
	}
}

func _Logic_SingleRecall0_HTTP_Handler(srv LogicHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in SingleRecallRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationLogicSingleRecall)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.SingleRecall(ctx, req.(*SingleRecallRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*RecallReplay)
		return ctx.Result(200, reply)
	}
}

func _Logic_GroupSend0_HTTP_Handler(srv LogicHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GroupSendRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationLogicGroupSend)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GroupSend(ctx, req.(*GroupSendRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*SendReplay)
		return ctx.Result(200, reply)
	}
}

func _Logic_GroupSendMention0_HTTP_Handler(srv LogicHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GroupSendMentionRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationLogicGroupSendMention)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GroupSendMention(ctx, req.(*GroupSendMentionRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*SendReplay)
		return ctx.Result(200, reply)
	}
}

func _Logic_GroupRecall0_HTTP_Handler(srv LogicHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GroupRecallRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationLogicGroupRecall)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GroupRecall(ctx, req.(*GroupRecallRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*RecallReplay)
		return ctx.Result(200, reply)
	}
}

func _Logic_RoomSend0_HTTP_Handler(srv LogicHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GroupSendRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationLogicRoomSend)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.RoomSend(ctx, req.(*GroupSendRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*SendReplay)
		return ctx.Result(200, reply)
	}
}

func _Logic_RoomBroadcast0_HTTP_Handler(srv LogicHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GroupSendRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationLogicRoomBroadcast)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.RoomBroadcast(ctx, req.(*GroupSendRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*SendReplay)
		return ctx.Result(200, reply)
	}
}

func _Logic_Login0_HTTP_Handler(srv LogicHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in LoginReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationLogicLogin)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Login(ctx, req.(*LoginReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*LoginReplay)
		return ctx.Result(200, reply)
	}
}

func _Logic_GetService0_HTTP_Handler(srv LogicHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetServiceReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationLogicGetService)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetService(ctx, req.(*GetServiceReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetServiceReplay)
		return ctx.Result(200, reply)
	}
}

type LogicHTTPClient interface {
	GetService(ctx context.Context, req *GetServiceReq, opts ...http.CallOption) (rsp *GetServiceReplay, err error)
	GroupRecall(ctx context.Context, req *GroupRecallRequest, opts ...http.CallOption) (rsp *RecallReplay, err error)
	GroupSend(ctx context.Context, req *GroupSendRequest, opts ...http.CallOption) (rsp *SendReplay, err error)
	GroupSendMention(ctx context.Context, req *GroupSendMentionRequest, opts ...http.CallOption) (rsp *SendReplay, err error)
	Login(ctx context.Context, req *LoginReq, opts ...http.CallOption) (rsp *LoginReplay, err error)
	RoomBroadcast(ctx context.Context, req *GroupSendRequest, opts ...http.CallOption) (rsp *SendReplay, err error)
	RoomSend(ctx context.Context, req *GroupSendRequest, opts ...http.CallOption) (rsp *SendReplay, err error)
	SingleRecall(ctx context.Context, req *SingleRecallRequest, opts ...http.CallOption) (rsp *RecallReplay, err error)
	SingleSend(ctx context.Context, req *SingleSendRequest, opts ...http.CallOption) (rsp *SendReplay, err error)
}

type LogicHTTPClientImpl struct {
	cc *http.Client
}

func NewLogicHTTPClient(client *http.Client) LogicHTTPClient {
	return &LogicHTTPClientImpl{client}
}

func (c *LogicHTTPClientImpl) GetService(ctx context.Context, in *GetServiceReq, opts ...http.CallOption) (*GetServiceReplay, error) {
	var out GetServiceReplay
	pattern := "v1/service"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationLogicGetService))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *LogicHTTPClientImpl) GroupRecall(ctx context.Context, in *GroupRecallRequest, opts ...http.CallOption) (*RecallReplay, error) {
	var out RecallReplay
	pattern := "v1/group/recall"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationLogicGroupRecall))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *LogicHTTPClientImpl) GroupSend(ctx context.Context, in *GroupSendRequest, opts ...http.CallOption) (*SendReplay, error) {
	var out SendReplay
	pattern := "v1/group/send"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationLogicGroupSend))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *LogicHTTPClientImpl) GroupSendMention(ctx context.Context, in *GroupSendMentionRequest, opts ...http.CallOption) (*SendReplay, error) {
	var out SendReplay
	pattern := "v1/group/send_mention"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationLogicGroupSendMention))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *LogicHTTPClientImpl) Login(ctx context.Context, in *LoginReq, opts ...http.CallOption) (*LoginReplay, error) {
	var out LoginReplay
	pattern := "v1/login"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationLogicLogin))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *LogicHTTPClientImpl) RoomBroadcast(ctx context.Context, in *GroupSendRequest, opts ...http.CallOption) (*SendReplay, error) {
	var out SendReplay
	pattern := "v1/room/broadcast"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationLogicRoomBroadcast))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *LogicHTTPClientImpl) RoomSend(ctx context.Context, in *GroupSendRequest, opts ...http.CallOption) (*SendReplay, error) {
	var out SendReplay
	pattern := "v1/room/send"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationLogicRoomSend))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *LogicHTTPClientImpl) SingleRecall(ctx context.Context, in *SingleRecallRequest, opts ...http.CallOption) (*RecallReplay, error) {
	var out RecallReplay
	pattern := "v1/single/recall"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationLogicSingleRecall))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *LogicHTTPClientImpl) SingleSend(ctx context.Context, in *SingleSendRequest, opts ...http.CallOption) (*SendReplay, error) {
	var out SendReplay
	pattern := "v1/single/send"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationLogicSingleSend))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
