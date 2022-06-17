// Copyright 2022 gqzcl <gqzcl@qq.com>. All rights reserved.
// Use of this source code is governed by a MIT style

package data

import (
	"context"

	"github.com/gomodule/redigo/redis"
)

// uid -> address
// address -> server
// 添加用户连接地址
func (d *Data) AddUserAddress(c context.Context, uid, address, server string) {
	conn := d.redis.Get()
	defer conn.Close()
	_, err := conn.Do("HSET", uid, address, server)
	if err != nil {
		return
	}
}

// 获取用户的连接地址
func (d *Data) GetUserAddress(c context.Context, uid string) (address map[string]string, err error) {
	conn := d.redis.Get()
	defer conn.Close()
	do, err := redis.StringMap(conn.Do("HGETALL", uid))
	if err != nil {
		return nil, err
	}
	return do, err
}

// DelUserAddress 删除用户的连接地址
func (d *Data) DelUserAddress(c context.Context, uid, address string) (success bool, err error) {
	conn := d.redis.Get()
	defer conn.Close()
	_, err = conn.Do("HDEL", uid, address)
	if err != nil {
		return false, err
	}
	return true, err

}

func (d *Data) GetAuth(c context.Context, uid, address string) (token string, err error) {
	conn := d.redis.Get()
	defer conn.Close()
	_, err = conn.Do("HGET", uid, token)
	return
}

func (d *Data) SetAuth(c context.Context, uid, address string) (err error) {
	conn := d.redis.Get()
	defer conn.Close()
	//_, err = conn.Do("HGET", uid, token)
	return
}
