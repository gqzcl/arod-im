// Copyright 2022 gqzcl <gqzcl@qq.com>. All rights reserved.
// Use of this source code is governed by a MIT style

package conf

import (
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/model"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"log"
	"sync"
)

var (
	namingClient naming_client.INamingClient
	nonce        = &sync.Once{}
)

func init() {
	GetNamingClient()
}

// GetNamingClient create naming client
func GetNamingClient() naming_client.INamingClient {
	if namingClient == nil {
		nonce.Do(func() {
			var err error
			namingClient, err = clients.NewNamingClient(
				vo.NacosClientParam{
					ClientConfig:  &cc,
					ServerConfigs: sc,
				},
			)
			if err != nil {
				fmt.Println("Naming client create failed")
			} else {
				fmt.Println("Naming client create success")
			}
		})
	}
	return namingClient
}

func RegisterInstance() error {
	// TODO 直接从配置文件导入
	success, err := namingClient.RegisterInstance(vo.RegisterInstanceParam{
		Ip:          "10.0.0.11",
		Port:        8848,
		ServiceName: "demo.go",
		Weight:      10,
		Enable:      true,
		Healthy:     true,
		Ephemeral:   true,
		Metadata:    map[string]string{"idc": "shanghai"},
		ClusterName: "cluster-a", // default value is DEFAULT
		GroupName:   "group-a",   // default value is DEFAULT_GROUP
	})
	if err != nil {
		return err
	}
	if success {
		fmt.Println("Register Instance successfully.")
	}
	return err
}

func Deregisterinstance() error {
	success, err := namingClient.DeregisterInstance(vo.DeregisterInstanceParam{
		Ip:          "10.0.0.11",
		Port:        8848,
		ServiceName: "demo.go",
		Ephemeral:   true,
		Cluster:     "cluster-a", // default value is DEFAULT
		GroupName:   "group-a",   // default value is DEFAULT_GROUP
	})
	if err != nil {
		return err
	}
	if success {
		fmt.Println("Deregister instance successfully")
	}
	return err
}

func GetAllInstances() ([]model.Instance, error) {
	instances, err := namingClient.SelectAllInstances(vo.SelectAllInstancesParam{
		ServiceName: "demo.go",
		GroupName:   "group-a",             // default value is DEFAULT_GROUP
		Clusters:    []string{"cluster-a"}, // default value is DEFAULT
	})
	if err != nil {
		return nil, err
	}
	return instances, err
}

func GetInstance() ([]model.Instance, error) {
	instance, err := namingClient.SelectInstances(vo.SelectInstancesParam{
		ServiceName: "demo.go",
		GroupName:   "group-a",             // default value is DEFAULT_GROUP
		Clusters:    []string{"cluster-a"}, // default value is DEFAULT
		HealthyOnly: true,
	})
	if err != nil {
		return nil, err
	}
	return instance, err
}

// GetOneHealthyInstance return one instance by WRR strategy for load balance
// And the instance should be health=true,enable=true and weight>0
func GetOneHealthyInstance() (*model.Instance, error) {
	instance, err := namingClient.SelectOneHealthyInstance(vo.SelectOneHealthInstanceParam{
		ServiceName: "demo.go",
		GroupName:   "group-a",             // default value is DEFAULT_GROUP
		Clusters:    []string{"cluster-a"}, // default value is DEFAULT
	})
	if err != nil {
		return nil, err
	}
	return instance, err
}

func SubscribeClient() error {
	err := namingClient.Subscribe(&vo.SubscribeParam{
		ServiceName: "demo.go",
		GroupName:   "group-a",             // default value is DEFAULT_GROUP
		Clusters:    []string{"cluster-a"}, // default value is DEFAULT
		SubscribeCallback: func(services []model.SubscribeService, err error) {
			log.Printf("\n\n callback return services:%s \n\n", services)
		},
	})
	return err
}

func UnSubscribeClient() error {
	err := namingClient.Unsubscribe(&vo.SubscribeParam{
		ServiceName: "demo.go",
		GroupName:   "group-a",             // default value is DEFAULT_GROUP
		Clusters:    []string{"cluster-a"}, // default value is DEFAULT
		SubscribeCallback: func(services []model.SubscribeService, err error) {
			log.Printf("\n\n callback return services:%s \n\n", services)
		},
	})
	return err
}
func GetAllServiceName() (model.ServiceList, error) {
	serviceInfos, err := namingClient.GetAllServicesInfo(vo.GetAllServiceInfoParam{
		NameSpace: "0e83cc81-9d8c-4bb8-a28a-ff703187543f",
		PageNo:    1,
		PageSize:  10,
	})
	return serviceInfos, err
}
