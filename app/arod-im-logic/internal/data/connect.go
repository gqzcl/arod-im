// Copyright 2022 gqzcl <gqzcl@qq.com>. All rights reserved.
// Use of this source code is governed by a MIT style

package data

import (
	"arod-im/app/arod-im-logic/internal/biz"
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type connectRepo struct {
	data *Data
	log  *log.Helper
}

func NewConnectRepo(data *Data, logger log.Logger) biz.ConnectRepo {
	return &connectRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (c *connectRepo) Connect(ctx context.Context, uid string, address string, server string) (err error) {
	c.data.AddUserAddress(ctx, uid, address, server)
	if err != nil {
		c.log.Infof("Connect set receive occur error: %v", err)
		return err
	}
	return nil
}
func (c *connectRepo) Disconnect(ctx context.Context, uid string, address string, server string) (err error) {
	success, err := c.data.DelUserAddress(ctx, uid, address)
	if !success {
		c.log.WithContext(ctx).Errorf("error in Disconnect() err:%v", err)
		return err
	}
	return nil
}
