package service

import (
	v1 "arod-im/api/logic/v1"
	"context"
)

// StoreConnect 发送连接信息到业务服务
func (s *ConnectorService) StoreConnect(uid, address string) (bool, error) {
start:
	client := s.discovery.GetClient()
	if client == nil {
		goto start
	}
	connect, err := client.Connect(context.Background(), &v1.ConnectReq{
		Server:  s.Address,
		Uid:     uid,
		Address: address,
		Token:   nil,
	})
	if err != nil {
		return false, err
	}
	return connect.Success, nil
}
