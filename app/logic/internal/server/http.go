// Copyright 2022 gqzcl <gqzcl@qq.com>. All rights reserved.
// Use of this source code is governed by a MIT style

package server

import (
	v1 "arod-im/api/logic/v1"
	"arod-im/app/logic/internal/conf"
	"arod-im/app/logic/internal/service"
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/gorilla/handlers"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
)

func SkipRouterMatch() selector.MatchFunc {
	whiteList := make(map[string]struct{})
	whiteList["/api.logic.v1.Logic/Login"] = struct{}{}
	return func(ctx context.Context, operation string) bool {
		if _, ok := whiteList[operation]; ok {
			return false
		}
		return true
	}
}

// Set global trace provider
func setTracerProvider(url string) (*tracesdk.TracerProvider, error) {
	// Create the Jaeger exporter
	exp, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(url)))
	if err != nil {
		return nil, err
	}

	tp := tracesdk.NewTracerProvider(
		tracesdk.WithSampler(tracesdk.ParentBased(tracesdk.TraceIDRatioBased(1.0))),
		// Always be sure to batch in production.
		tracesdk.WithBatcher(exp),
		// Record information about this application in an Resource.
		tracesdk.WithResource(resource.NewSchemaless(
			semconv.ServiceNameKey.String("logic"),
			attribute.String("env", "dev"),
			attribute.Int64("ID", 1),
		)),
		// tracesdk.WithResource(resource.NewWithAttributes(
		// 	semconv.SchemaURL,
		// 	semconv.ServiceNameKey.String(Name),
		// 	semconv.ServiceVersionKey.String(Version),
		// 	attribute.String("environment", "dev"),
		// )),
	)

	//otel.SetTracerProvider(tp)
	return tp, nil
}

// NewHTTPServer new a HTTP server.
func NewHTTPServer(c *conf.Server, message *service.MessageService, logger log.Logger) *http.Server {
	url := "http://127.0.0.1:5000/api/traces"
	tp, err := setTracerProvider(url)
	if err != nil {
		panic(err)
	}
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			selector.Server(JWTAuth()).Match(SkipRouterMatch()).Build(),
			tracing.Server(
				tracing.WithTracerProvider(tp),
				tracing.WithPropagator(
					propagation.NewCompositeTextMapPropagator(propagation.Baggage{}, propagation.TraceContext{}),
				),
			),
		),
		http.Filter(handlers.CORS(
			handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
			handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "HEAD", "OPTIONS"}),
			handlers.AllowedOrigins([]string{"*"}),
		)),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	v1.RegisterLogicHTTPServer(srv, message)
	return srv
}
