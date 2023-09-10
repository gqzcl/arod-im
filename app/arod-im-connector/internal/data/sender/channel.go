// Copyright 2022 gqzcl <gqzcl@qq.com>. All rights reserved.
// Use of this source code is governed by a MIT style

package sender

import (
	jobV1 "arod-im/api/job/v1"
	"encoding/json"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"github.com/panjf2000/gnet/v2"
)

type Channel struct {
	Room *Room
	Next *Channel
	Prev *Channel

	conn gnet.Conn

	Uid  string
	Addr string
}

func NewChannel(conn gnet.Conn) *Channel {
	return &Channel{
		conn: conn,
	}
}

func (r *Channel) Push(msg []*jobV1.MsgBody) error {
	m, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	//fmt.Println("消息json序列号完成", m)
	err = wsutil.WriteServerMessage(r.conn, ws.OpText, m)
	if err != nil {
		return err
	}
	return nil
}
