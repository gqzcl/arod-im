package service

import (
	"arod-im/api/logic/v1"
	"arod-im/app/comet/internal/conf"
	"arod-im/app/comet/internal/pkg"
	"context"
	"math/rand"
	"time"

	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/google/wire"
	"github.com/zhenjl/cityhash"
	"go.opentelemetry.io/otel/sdk/trace"
)

const (
	minServerHeartbeat = time.Minute * 10
	maxServerHeartbeat = time.Minute * 30
)

var ProciderSet = wire.NewSet(NewLogicClient, NewServer)

// 集中管理Comet服务器资源
type Server struct {
	round      *pkg.Round
	buckets    []*pkg.Bucket
	bucketSize uint32
	serverID   string
	lc         logic.LogicClient
}

func NewServer(size uint32) *Server {
	s := &Server{
		round:      pkg.NewRound(&conf.TCP{}),
		bucketSize: size,
	}
	s.buckets = make([]*pkg.Bucket, size)
	for i := 0; i < len(s.buckets); i++ {
		s.buckets[i] = pkg.NewBucket(&conf.Bucket{})
	}
	go s.onlineproc()
	return s
}

func NewLogicClient(r registry.Discovery, trace *trace.TracerProvider) logic.LogicClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovert://arod-im.logic"),
		grpc.WithDiscovery(r),
		grpc.WithMiddleware(
			tracing.Client(tracing.WithTracerProvider(trace)),
			recovery.Recovery(),
		),
	)
	if err != nil {
		panic(err)
	}
	return logic.NewLogicClient(conn)
}
func (s *Server) onlineproc() {
	// TODO 统计房间人数
}

func (s *Server) ListBucket() []*pkg.Bucket {
	return s.buckets
}

// 通过subkey获取对应的bucket
func (s *Server) GetBucket(subKey string) *pkg.Bucket {
	idx := cityhash.CityHash32([]byte(subKey), uint32(len(subKey))) % s.bucketSize
	return s.buckets[idx]
}

// 生成介于10分钟到30分钟之间的随机的服务器心跳时间
func (s *Server) RandServerHearbeat() time.Duration {
	return (minServerHeartbeat + time.Duration(rand.Int63n(int64(maxServerHeartbeat-minServerHeartbeat))))
}
