package server

import (
	"arod-im/app/job/internal/conf"
	"github.com/go-kratos/kratos/contrib/registry/nacos/v2"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

func NewNacosRegister(c *conf.Server) *nacos.Registry {
	cc := constant.ClientConfig{
		NamespaceId:         c.Register.NamespaceId,
		AccessKey:           c.Register.AccessKey,
		SecretKey:           c.Register.SecretKey,
		TimeoutMs:           c.Register.TimeoutMs,
		NotLoadCacheAtStart: true,
		LogDir:              "/tmp/nacos/log",
		CacheDir:            "/tmp/nacos/cache",
		BeatInterval:        30000,
		LogLevel:            c.Register.LogLevel,
	}
	sc := []constant.ServerConfig{
		*constant.NewServerConfig(c.Register.Address, c.Register.Port),
	}
	namingClient, err := clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  &cc,
			ServerConfigs: sc,
		},
	)
	if err != nil {
		panic(err)
	}

	//namingClient.GetService()
	return nacos.New(namingClient, nacos.WithGroup("arod-im"))
}
