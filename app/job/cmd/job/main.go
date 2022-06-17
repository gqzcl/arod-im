// Copyright 2022 gqzcl <gqzcl@qq.com>. All rights reserved.
// Use of this source code is governed by a MIT style

package main

import (
	"arod-im/app/job/internal/conf"
	nacosConfig "arod-im/pkg/config"
	"arod-im/pkg/transport/kafka"
	"flag"
	"os"
	"time"

	"github.com/go-kratos/kratos/contrib/registry/nacos/v2"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/grpc"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	Name string
	// Version is the version of the compiled software.
	Version string
	// flagconf is the config flag.
	flagconf string

	id, _ = os.Hostname()
)

func init() {
	flag.StringVar(&flagconf, "conf", "../../configs", "config path, eg: -conf config.yaml")
}

func newApp(logger log.Logger, gs *grpc.Server, ks *kafka.Server, r *nacos.Registry) *kratos.App {
	return kratos.New(
		kratos.ID(id),
		kratos.Name("arod-im-job"),
		kratos.Version("v0.1.0"),
		kratos.Metadata(map[string]string{"preserved.heart.beat.interval": "300000"}),
		kratos.Logger(logger),
		kratos.Server(
			gs,
			ks,
		),
		kratos.Registrar(r),
		kratos.RegistrarTimeout(time.Duration(5000)),
	)
}

func main() {
	flag.Parse()
	logger := log.With(log.NewStdLogger(os.Stdout),
		"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
		"service.id", id,
		"service.name", Name,
		"service.version", Version,
		"trace.id", tracing.TraceID(),
		"span.id", tracing.SpanID(),
	)
	c := config.New(
		config.WithSource(
			// 后面的会覆盖前面的
			nacosConfig.NewLocalConfigSource(flagconf),
			nacosConfig.NewNacosConfigSource("localhost", 8848, "arod-im-job-1.yaml"),
		),
	)
	defer c.Close()

	if err := c.Load(); err != nil {
		panic(err)
	}

	var bc conf.Bootstrap
	if err := c.Scan(&bc); err != nil {
		panic(err)
	}

	app, cleanup, err := wireApp(bc.Server, bc.Data, logger)
	if err != nil {
		panic(err)
	}
	defer cleanup()
	// start and wait for stop signal
	if err := app.Run(); err != nil {
		panic(err)
	}
}
