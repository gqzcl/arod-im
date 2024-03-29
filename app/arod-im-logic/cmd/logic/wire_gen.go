// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"arod-im/app/arod-im-logic/internal/biz"
	"arod-im/app/arod-im-logic/internal/conf"
	"arod-im/app/arod-im-logic/internal/data"
	"arod-im/app/arod-im-logic/internal/server"
	"arod-im/app/arod-im-logic/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, confData *conf.Data, logger log.Logger) (*kratos.App, func(), error) {
	syncProducer := data.NewKafkaPub(confData)
	pool := data.NewRedis(confData)
	dataData, cleanup, err := data.NewData(confData, logger, syncProducer, pool)
	if err != nil {
		return nil, nil, err
	}
	singleRepo := data.NewSingleRepo(dataData, logger)
	singleUsecase := biz.NewSingleUsecase(singleRepo, logger)
	groupRepo := data.NewGroupRepo(dataData, logger)
	groupUsecase := biz.NewGroupUsecase(groupRepo, logger)
	roomRepo := data.NewRoomRepo(dataData, logger)
	roomUsecase := biz.NewRoomUsecase(roomRepo, logger)
	connectRepo := data.NewConnectRepo(dataData, logger)
	connectUsecase := biz.NewConnectUsecase(connectRepo, logger)
	loginRepo := data.NewLoginRepo(dataData, logger)
	loginUsecase := biz.NewLoginUsecase(loginRepo, logger)
	messageService := service.NewMessageService(singleUsecase, groupUsecase, roomUsecase, connectUsecase, loginUsecase, logger)
	httpServer := server.NewHTTPServer(confServer, messageService, logger)
	grpcServer := server.NewGRPCServer(confServer, messageService, logger)
	registry := server.NewNacosRegister(confServer, dataData)
	app := newApp(logger, httpServer, grpcServer, registry)
	return app, func() {
		cleanup()
	}, nil
}
