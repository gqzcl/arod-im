// Copyright 2022 gqzcl <gqzcl@qq.com>. All rights reserved.
// Use of this source code is governed by a MIT style

package biz

import "github.com/go-kratos/kratos/v2/log"

type RoomRepo interface {
}

// MessageUsecase  is a Message usecase.
type RoomUsecase struct {
	room RoomRepo
	log  *log.Helper
}

// NewMessageUsecase  new a Message usecase.
func NewRoomUsecase(room RoomRepo, logger log.Logger) *RoomUsecase {
	return &RoomUsecase{room: room, log: log.NewHelper(logger)}
}
