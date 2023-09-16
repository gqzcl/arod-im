// Copyright 2022 gqzcl <gqzcl@qq.com>. All rights reserved.
// Use of this source code is governed by a MIT style

package ips

import (
	"net"
	"strings"
)

// InternalIP return internal ip.
func InternalIP() string {
	networks, err := net.Interfaces()
	if err != nil {
		return ""
	}
	for _, network := range networks {
		if network.Flags&net.FlagUp != 0 && !strings.HasPrefix(network.Name, "lo") {
			addrs, err := network.Addrs()
			if err != nil {
				continue
			}
			for _, addr := range addrs {
				if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
					if ipnet.IP.To4() != nil || ipnet.IP.To16() != nil {
						return ipnet.IP.String()
					}
				}
			}
		}
	}
	return ""
}
