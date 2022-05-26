// Copyright 2022 gqzcl <gqzcl@qq.com>. All rights reserved.
// Use of this source code is governed by a MIT style

package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

// ConnectRepo is a Connection repo.
type ConnectRepo interface {
	//Push(ctx context.Context, sessionId string, msg []*v1.SingleSendRequest_MsgBody) (err error)
	Connect(ctx context.Context, uid string, address string, server string) (err error)
	Disconnect(ctx context.Context, uid string, address string, server string) (err error)
}

// ConnectUsecase is a Connection  use case.
type ConnectUsecase struct {
	connect ConnectRepo
	log     *log.Helper
}

// NewConnectUsecase new a Connection use case.
func NewConnectUsecase(connect ConnectRepo, logger log.Logger) *ConnectUsecase {
	return &ConnectUsecase{connect: connect, log: log.NewHelper(logger)}
}

func (cc *ConnectUsecase) Connect(ctx context.Context, uid string, address string, server string) {
	cc.log.WithContext(ctx).Infof("uid:%s,address:%s,server:%s", uid, address, server)
	cc.connect.Connect(ctx, uid, address, server)
}
func (cc *ConnectUsecase) Disconnect(ctx context.Context, uid string, address string, server string) {
	cc.log.WithContext(ctx).Infof("uid:%s,address:%s,server:%s", uid, address, server)
	cc.connect.Disconnect(ctx, uid, address, server)
}
