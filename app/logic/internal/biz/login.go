package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type LoginRepo interface {
	GetServiceAddress(ctx context.Context) string
}

type LoginUsecase struct {
	login LoginRepo
	log   *log.Helper
}

func NewLoginUsecase(login LoginRepo, logger log.Logger) *LoginUsecase {
	return &LoginUsecase{login: login, log: log.NewHelper(logger)}
}

// GetServiceAddress 获取连接服务地址
func (lc *LoginUsecase) GetServiceAddress(ctx context.Context) string {
	return lc.login.GetServiceAddress(ctx)
}
