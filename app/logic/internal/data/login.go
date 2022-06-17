package data

import (
	"arod-im/app/logic/internal/biz"

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

func (l *loginRepo) CheckAuth(username, password string) bool {
	l.data.redis.Get()
	return true
}
