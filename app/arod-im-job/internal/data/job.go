// Copyright 2022 gqzcl <gqzcl@qq.com>. All rights reserved.
// Use of this source code is governed by a MIT style

package data

import (
	v1 "arod-im/api/connector/v1"
	jobV1 "arod-im/api/job/v1"
	"arod-im/app/arod-im-job/internal/biz"
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

var _ biz.JobRepo = (*jobRepo)(nil)

type jobRepo struct {
	data *Data
	log  *log.Helper
}

// NewJobRepo NewGreeterRepo .
func NewJobRepo(data *Data, logger log.Logger) biz.JobRepo {
	return &jobRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (j *jobRepo) SingleSend(ctx context.Context, address, server, senderId, seq string, msg []*jobV1.MsgBody) error {
	j.log.WithContext(ctx).Debug("开始发送消息")

	//connector:=j.data.discovery.GetClient(server)
	client := j.data.discovery.GetClient(server)
	if client != nil {
		// TODO 为client加入环形队列，进行缓存
		sendReply, err := client.SingleSend(ctx, &v1.SingleSendReq{
			Address:  address,
			SenderId: senderId,
			Seq:      seq,
			Msg:      msg,
		})
		if err != nil {
			j.log.WithContext(ctx).Error(err)
		}
		j.log.WithContext(ctx).Infof("Response received: %v", sendReply)
	} else {
		j.log.WithContext(ctx).Debugf("连接服务:%s 未上线", server)
	}
	return nil
}
