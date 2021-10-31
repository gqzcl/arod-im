package main

import (
	"flag"

	"arod-im/app/comet/internal/conf"
	"arod-im/app/comet/internal/logger"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
)

const (
	Version = "0.1.0"
	Name    = "goim.logic"
)

var flagconf string

func init() {
	flag.StringVar(&flagconf, "conf", "../configs", "config path, eg: -conf config.yaml")
}

func newApp(logger log.Logger, rr registry.Registrar) *kratos.App {
	return kratos.New(
		kratos.Name(Name),
		kratos.Version(Version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Registrar(rr),
	)
}
func main() {
	flag.Parse()
	// 加载日志
	// logger := log.With(log.NewStdLogger(os.Stdout),
	// 	"service.name", Name,
	// 	"service.version", Version,
	// 	"ts", log.DefaultTimestamp,
	// 	"caller", log.DefaultCaller,
	// )
	logger := log.With(logger.NewZapLogger(),
		"service.name", Name,
		"service.version", Version,
		"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
	)

	// 加载配置文件
	c := config.New(
		config.WithSource(
			file.NewSource(flagconf),
		),
	)
	if err := c.Load(); err != nil {
		panic(err)
	}

	// var bc conf.Bootstrap
	// if err := c.Scan(&bc); err != nil {
	// 	panic(err)
	// }

	// 服务注册
	var register conf.Consul
	if err := c.Scan(&register); err != nil {
		panic(err)
	}

	// // 指标监控
	// exp, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint("123")))
	// if err != nil {
	// 	panic(err)
	// }

	// // 链路追踪
	// tp := tracesdk.NewTracerProvider(
	// 	tracesdk.WithBatcher(exp),
	// 	tracesdk.WithResource(resource.NewSchemaless(
	// 		semconv.ServiceNameKey.String(Name),
	// 	)),
	// )

	app, cleanup, err := initApp(&register, logger)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	if err := app.Run(); err != nil {
		panic(err)
	}
}
