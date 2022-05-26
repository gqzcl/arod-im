// Copyright 2022 gqzcl <gqzcl@qq.com>. All rights reserved.
// Use of this source code is governed by a MIT style

package ips

import (
	"fmt"
	"testing"
)

func TestIP(t *testing.T) {
	ip := InternalIP()
	fmt.Println(ip)
	if ip == "" {
		t.FailNow()
	}
}
