// Copyright 2022 gqzcl <gqzcl@qq.com>. All rights reserved.
// Use of this source code is governed by a MIT style

package biz

import "github.com/go-kratos/kratos/v2/log"

type GroupRepo interface {
}

// MessageUsecase  is a Message usecase.
type GroupUsecase struct {
	group GroupRepo
	log   *log.Helper
}

// NewMessageUsecase  new a Message usecase.
func NewGroupUsecase(group GroupRepo, logger log.Logger) *GroupUsecase {
	return &GroupUsecase{group: group, log: log.NewHelper(logger)}
}
