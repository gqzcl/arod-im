package main

import (
	logic "arod-im/app/logic/internal"
	"arod-im/app/logic/internal/conf"
	"arod-im/app/logic/internal/http"
	"flag"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/golang/glog"
	"github.com/hashicorp/consul/api"
)

const (
	Version = "0.1.0"
	Name    = "goim.logic"
)

func init() {
	//	解析命令行标志
	flag.Parse()
	// 初始化配置文件
	if err := conf.Init(); err != nil {
		panic(err)
	}
	// 打印版本信息和环境信息
	glog.Infof("goim-logic [version: %s env: %+v] start", Version, conf.Conf.Env)
}

func newApp(logger log.Logger, gs *grpc.Server, rr registry.Registrar) *kratos.App {
	return kratos.New(
		kratos.Name(Name),
		kratos.Version(Version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Server(
			gs,
		),
		kratos.Registrar(rr),
	)
}

func main() {
	// grpc register
	client, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		panic(err)
	}

	// dis := naming.New(conf.Conf.Discovery)
	// resolver.Register(dis)

	// logic
	srv := logic.New(conf.Conf)
	httpSrv := http.New(conf.Conf.HTTPServer, srv)
	rpcSrv := grpc.New(conf.Conf.RPCServer, srv)

	// cancel := register(dis, srv)

	// // signal
	// c := make(chan os.Signal, 1)
	// signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	// for {
	// 	s := <-c
	// 	log.Infof("goim-logic get a signal %s", s.String())
	// 	switch s {
	// 	case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
	// 		if cancel != nil {
	// 			cancel()
	// 		}
	// 		srv.Close()
	// 		httpSrv.Close()
	// 		rpcSrv.GracefulStop()
	// 		log.Infof("goim-logic [version: %s] exit", ver)
	// 		log.Flush()
	// 		return
	// 	case syscall.SIGHUP:
	// 	default:
	// 		return
	// 	}
	// }
}

// func register(dis *naming.Discovery, srv *logic.Logic) context.CancelFunc {
// 	env := conf.Conf.Env
// 	addr := ips.InternalIP()
// 	_, port, _ := net.SplitHostPort(conf.Conf.RPCServer.Addr)
// 	ins := &naming.Instance{
// 		Region:   env.Region,
// 		Zone:     env.Zone,
// 		Env:      env.DeployEnv,
// 		Hostname: env.Host,
// 		AppID:    appid,
// 		Addrs: []string{
// 			"grpc://" + addr + ":" + port,
// 		},
// 		Metadata: map[string]string{
// 			model.MetaWeight: strconv.FormatInt(env.Weight, 10),
// 		},
// 	}
// 	cancel, err := dis.Register(ins)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return cancel
// }
