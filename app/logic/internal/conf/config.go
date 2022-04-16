package conf

import (
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"sync"
)

//var endpoint = "101.43.63.229"
var namespaceId = ""
var accessKey = "nacos"
var secretKey = "nacos"
var dataId = "com.gqzcl.arod-im.yaml"
var group = "CONNECTOR"

type Clients struct {
	configClient config_client.IConfigClient
}

var (
	configClient config_client.IConfigClient
	once         = &sync.Once{}
	cc           constant.ClientConfig
	sc           []constant.ServerConfig
)

func init() {
	cc = constant.ClientConfig{
		//Endpoint:    endpoint + "8848",
		NamespaceId: namespaceId,
		AccessKey:   accessKey,
		SecretKey:   secretKey,
		TimeoutMs:   5 * 1000,
	}
	sc = []constant.ServerConfig{
		{
			IpAddr: "101.43.63.229",
			Port:   8848,
		},
	}
	GetConfigClient()
}

// GetConfigClient create config client
func GetConfigClient() config_client.IConfigClient {
	if configClient == nil {
		once.Do(func() {
			var err error
			//cc = &Clients{}
			configClient, err = clients.NewConfigClient(
				vo.NacosClientParam{
					ClientConfig:  &cc,
					ServerConfigs: sc,
				},
			)
			if err != nil {
				fmt.Println("Config client create failed")
			} else {
				fmt.Println("Config client create success")
			}
		})
	}
	return configClient
}

//func GetConfig() (*Config, error) {
//	//cc := GetConfigClient()
//	// TODO 可以考虑改造成单例
//	// 获取配置
//	content, err := configClient.GetConfig(vo.ConfigParam{
//		DataId: dataId,
//		Group:  group,
//	})
//	c := new(Config)
//	var configContent string
//	flag.StringVar(&configContent, "conf", content, "-config")
//	err = yaml.Unmarshal([]byte(configContent), c)
//	// fmt.Println(c.TCP)
//	if err != nil {
//		fmt.Println("Unmarshal failed in GetConfig()")
//		return nil, err
//	}
//	return c, nil
//}

func ConfigPublish(content string) error {
	// cc := GetConfigClient()
	// 发布配置
	success, err := configClient.PublishConfig(vo.ConfigParam{
		DataId:  dataId,
		Group:   group,
		Content: content,
	})
	if err != nil {
		return err
	}
	if success {
		fmt.Println("Publish config successfully.")
	}
	return nil
}
