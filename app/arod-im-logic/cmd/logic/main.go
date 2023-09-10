// Copyright 2022 gqzcl <gqzcl@qq.com>. All rights reserved.
// Use of this source code is governed by a MIT style

package main

import (
	"arod-im/app/arod-im-logic/internal/conf"
	nacosConfig "arod-im/pkg/nacos/config"
	"flag"
	"os"

	"github.com/go-kratos/kratos/contrib/registry/nacos/v2"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	Name = "logic"
	// Version is the version of the compiled software.
	Version = "0.1.0"
	// flagconf is the config flag.
	flagconf string

	id, _ = os.Hostname()
)

func init() {
	flag.StringVar(&flagconf, "conf", "../../configs/config.yaml", "config path, eg: -conf config.yaml")
}

func newApp(logger log.Logger, hs *http.Server, gs *grpc.Server, r *nacos.Registry) *kratos.App {
	return kratos.New(
		kratos.ID(id),
		kratos.Name("arod-im-logic"),
		kratos.Version("v0.1.0"),
		// 心跳时间30s
		kratos.Metadata(map[string]string{"preserved.heart.beat.interval": "30000"}),
		kratos.Logger(logger),
		kratos.Server(
			hs,
			gs,
		),
		kratos.Registrar(r),
	)
}

// setTracerProvider Set global trace provider
func setTracerProvider(url string) error {
	// Create the Jaeger exporter
	exp, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(url)))
	if err != nil {
		return err
	}

	tp := tracesdk.NewTracerProvider(
		tracesdk.WithSampler(tracesdk.ParentBased(tracesdk.TraceIDRatioBased(1.0))),
		// Always be sure to batch in production.
		tracesdk.WithBatcher(exp),
		// Record information about this application in an Resource.
		tracesdk.WithResource(resource.NewSchemaless(
			semconv.ServiceNameKey.String(Name),
			semconv.ServiceVersionKey.String(Version),
			attribute.String("environment", "dev"),
			attribute.Int64("ID", 1),
		)),
	)

	otel.SetTracerProvider(tp)
	return nil
}

func main() {
	flag.Parse()
	// 日志格式
	logger := log.With(log.NewStdLogger(os.Stdout),
		"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
		"service.id", id,
		"service.name", Name,
		"service.version", Version,
		"trace.id", tracing.TraceID(),
		"span.id", tracing.SpanID(),
	)

	// config
	c := config.New(
		config.WithSource(
			// 后面的会覆盖前面的
			nacosConfig.NewLocalConfigSource(flagconf),
			nacosConfig.NewNacosConfigSource("localhost", 8848, "arod-im-logic-1.yaml"),
		),
	)
	defer c.Close()

	if err := c.Load(); err != nil {
		panic(err)
	}

	// 解析config
	var bc conf.Bootstrap
	if err := c.Scan(&bc); err != nil {
		panic(err)
	}

	// trace
	// url := "http://localhost:14268/api/traces"
	url := bc.Server.JaegerAddr
	err := setTracerProvider(url)
	if err != nil {
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
