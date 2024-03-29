// Copyright 2022 gqzcl <gqzcl@qq.com>. All rights reserved.
// Use of this source code is governed by a MIT style

package service

import (
	v1 "arod-im/api/logic/v1"
	"arod-im/app/arod-im-logic/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewMessageService)

// MessageService is a message service.
type MessageService struct {
	v1.UnimplementedLogicServer

	sc *biz.SingleUsecase
	gc *biz.GroupUsecase
	rc *biz.RoomUsecase
	cc *biz.ConnectUsecase
	lc *biz.LoginUsecase

	log *log.Helper
}

// NewMessageService new a message service.
func NewMessageService(sc *biz.SingleUsecase, gc *biz.GroupUsecase, rc *biz.RoomUsecase, cc *biz.ConnectUsecase, lc *biz.LoginUsecase, logger log.Logger) *MessageService {
	return &MessageService{
		sc:  sc,
		gc:  gc,
		rc:  rc,
		cc:  cc,
		lc:  lc,
		log: log.NewHelper(log.With(logger, "module", "logic")),
	}
}
