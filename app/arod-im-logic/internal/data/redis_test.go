// Copyright 2022 gqzcl <gqzcl@qq.com>. All rights reserved.
// Use of this source code is governed by a MIT style

package data

import (
	"context"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"testing"
	"time"
)

func TestData_GetUserAddress(t *testing.T) {
	r := &redis.Pool{
		MaxIdle:     int(100),
		MaxActive:   int(10),
		IdleTimeout: time.Second,
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial("tcp", "---",
				redis.DialConnectTimeout(time.Second*5),
				redis.DialReadTimeout(time.Second),
				redis.DialWriteTimeout(time.Second),
				redis.DialPassword("---"),
			)
			if err != nil {
				panic(err)
			}
			return conn, nil
		},
	}
	d := Data{redis: r}
	address, err := d.GetUserAddress(context.Background(), "100001")
	if err != nil {
		return
	}
	fmt.Println(address)
}
