// Copyright 2022 gqzcl <gqzcl@qq.com>. All rights reserved.
// Use of this source code is governed by a MIT style

package bytes

import (
	"fmt"
	"github.com/tx7do/kratos-transport/codec"
	"io"
)

type Codec struct {
	Conn io.ReadWriteCloser
}

type Frame struct {
	Data []byte
}

func (c *Codec) ReadHeader(m *codec.Message, t codec.MessageType) error {
	return nil
}

func (c *Codec) ReadBody(b interface{}) error {
	// read bytes
	buf, err := io.ReadAll(c.Conn)
	if err != nil {
		return err
	}

	switch v := b.(type) {
	case *[]byte:
		*v = buf
	case *Frame:
		v.Data = buf
	default:
		return fmt.Errorf("failed to read body: %v is not type of *[]byte", b)
	}

	return nil
}

func (c *Codec) Write(m *codec.Message, b interface{}) error {
	var v []byte
	switch vb := b.(type) {
	case *Frame:
		v = vb.Data
	case *[]byte:
		v = *vb
	case []byte:
		v = vb
	default:
		return fmt.Errorf("failed to write: %v is not type of *[]byte or []byte", b)
	}
	_, err := c.Conn.Write(v)
	return err
}

func (c *Codec) Close() error {
	return c.Conn.Close()
}

func (c *Codec) Name() string {
	return "bytes"
}

func NewCodec(c io.ReadWriteCloser) codec.Codec {
	return &Codec{
		Conn: c,
	}
}
