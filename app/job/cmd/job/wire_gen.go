// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

import (
	"arod-im/app/job/internal/biz"
	"arod-im/app/job/internal/conf"
	"arod-im/app/job/internal/data"
	"arod-im/app/job/internal/server"
	"arod-im/app/job/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, confData *conf.Data, logger log.Logger) (*kratos.App, func(), error) {
	dataData, cleanup, err := data.NewData(confData, logger)
	if err != nil {
		return nil, nil, err
	}
	jobRepo := data.NewJobRepo(dataData, logger)
	jobUsecase := biz.NewJobUsecase(jobRepo, logger)
	jobService := service.NewJobService(jobUsecase, logger)
	grpcServer := server.NewGRPCServer(confServer, jobService, logger)
	kafkaServer := server.NewKafkaServer(confServer, jobService)
	registry := server.NewNacosRegister(confServer, dataData)
	app := newApp(logger, grpcServer, kafkaServer, registry)
	return app, func() {
		cleanup()
	}, nil
}
