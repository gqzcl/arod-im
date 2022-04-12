package biz

import "github.com/go-kratos/kratos/v2/log"

type GroupRepo interface {
}

// MessageUsecase  is a Message usecase.
type GroupUsecase struct {
	group GroupRepo
	log   *log.Helper
}

// NewMessageUsecase  new a Message usecase.
func NewGroupUsecase(group GroupRepo, logger log.Logger) *GroupUsecase {
	return &GroupUsecase{group: group, log: log.NewHelper(logger)}
}
