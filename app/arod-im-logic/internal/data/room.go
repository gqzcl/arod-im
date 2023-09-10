// Copyright 2022 gqzcl <gqzcl@qq.com>. All rights reserved.
// Use of this source code is governed by a MIT style

package data

import (
	"arod-im/app/arod-im-logic/internal/biz"
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type roomRepo struct {
	data *Data
	log  *log.Helper
}

// NewGreeterRepo .
func NewRoomRepo(data *Data, logger log.Logger) biz.RoomRepo {
	return &roomRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *roomRepo) Send(ctx context.Context, g *biz.Greeter) (*biz.Greeter, error) {
	return g, nil
}

func (r *roomRepo) Update(ctx context.Context, g *biz.Greeter) (*biz.Greeter, error) {
	return g, nil
}

func (r *roomRepo) FindByID(context.Context, int64) (*biz.Greeter, error) {
	return nil, nil
}

func (r *roomRepo) ListByHello(context.Context, string) ([]*biz.Greeter, error) {
	return nil, nil
}

func (r *roomRepo) ListAll(context.Context) ([]*biz.Greeter, error) {
	return nil, nil
}
