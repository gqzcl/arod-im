package data

import (
	"arod-im/app/arod-im-logic/internal/biz"
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type Auth struct {
	ID       int    `gorm:"primary_key" json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type loginRepo struct {
	data *Data
	log  *log.Helper
}

func NewLoginRepo(data *Data, logger log.Logger) biz.LoginRepo {
	return &loginRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

// func (l *loginRepo) CheckAuth(username, password string) bool {
// 	l.data.redis.Get()
// 	return true
// }

// GetServiceAddress 获取连接服务地址
func (l *loginRepo) GetServiceAddress(ctx context.Context) string {
	serviceAddress := l.data.discovery.GetClient()
	return serviceAddress
}
