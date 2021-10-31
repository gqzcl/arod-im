package server

import (
	"arod-im/app/comet/internal/conf"

	consul "github.com/go-kratos/consul/registry"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/google/wire"
	consulAPI "github.com/hashicorp/consul/api"
)

// ProviderSet is server providers.
var ProviderSet = wire.NewSet(NewRegistrar)

func NewRegistrar(conf *conf.Consul) registry.Registrar {
	c := consulAPI.DefaultConfig()
	c.Address = conf.Address
	c.Scheme = conf.Scheme
	cli, err := consulAPI.NewClient(c)
	if err != nil {
		panic(err)
	}
	r := consul.New(cli, consul.WithHealthCheck(false))
	return r
}
