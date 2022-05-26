// Copyright 2022 gqzcl <gqzcl@qq.com>. All rights reserved.
// Use of this source code is governed by a MIT style

package data

import (
	jobV1 "arod-im/api/job/v1"
	"arod-im/app/connector/internal/biz"
	"arod-im/app/connector/internal/data/sender"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/panjf2000/gnet/v2"
)

type bucketRepo struct {
	data *Data
	log  *log.Helper
}

// NewBucketRepo NewBucketRepo .
func NewBucketRepo(data *Data, logger log.Logger) biz.BucketRepo {
	return &bucketRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (b *bucketRepo) RemoveCh(connectId string) {
	if room := b.data.channel[connectId].Room; room != nil {
		room.DelCh(b.data.channel[connectId])
	}
	delete(b.data.channel, connectId)
}

func (b *bucketRepo) AddCh(address string, c gnet.Conn) {
	// TODO 加锁
	b.log.Debug("成功添加channelwith", address)
	b.data.channel[address] = sender.NewChannel(c)
}

func (b *bucketRepo) SingleSend(address string, msg []*jobV1.MsgBody) {
	if c, ok := b.data.channel[address]; ok {
		err := c.Push(msg)
		if err != nil {
			b.log.Error(err, "in SingleSend in bucket")
		}
		b.log.Debug("成功发送消息")
		return
	}
	b.log.Warnf("消息发送失败:", address)
	//b.data.channel[address].Push(msg)
}

// PutChToRoom 当新连接的参数中带有roomID，则将他放到room中
func (b *bucketRepo) PutChToRoom(uid, roomId string) {
	b.data.room[roomId] = nil
}

// NewChannel 新建连接时创建一个channel
//func (b *bucketRepo) NewChannel(uid, connectId string) {
//	b.data.channel[uid] = sender.NewChannel(connectId)
//
//}
