package biz

import "github.com/go-kratos/kratos/v2/log"

type LoginRepo interface {
}

type LoginUsecase struct {
	login LoginRepo
	log   *log.Helper
}

func NewLoginUsecase(login LoginRepo, logger log.Logger) *LoginUsecase {
	return &LoginUsecase{login: login, log: log.NewHelper(logger)}
}
