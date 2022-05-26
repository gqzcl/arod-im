// Copyright 2022 gqzcl <gqzcl@qq.com>. All rights reserved.
// Use of this source code is governed by a MIT style

package kafka

var _ Consumer = (*Message)(nil)

type Message struct {
	Brokers   []string
	Topic     string
	partition int64
	h         Handler
}

func (c *Message) Key() string {
	return ""
}

func (c *Message) Value() []byte {
	return nil
}

func (c *Message) Header() map[string]string {
	return nil
}

func (c *Message) Ack() error {
	return nil
}

func (c *Message) Nack() error {
	return nil
}
