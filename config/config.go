package config

import (
	"github.com/dubbogo/gost/encoding/yaml"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

var DefaultConfig = &Config{}

func Init(nacosNamespace, nacosAddr string, port uint64, nacosUserName, nacosPwd string) {
	sc := []constant.ServerConfig{
		{
			IpAddr: nacosAddr,
			Port:   port,
		},
	}

	cc := constant.ClientConfig{
		NamespaceId:         nacosNamespace,
		NotLoadCacheAtStart: true,
		LogDir:              "./log",
		CacheDir:            "./cache",
		LogLevel:            "debug",
		Username:            nacosUserName,
		Password:            nacosPwd,
	}

	// a more graceful way to create config client
	client, err := clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  &cc,
			ServerConfigs: sc,
		},
	)

	if err != nil {
		panic(err)
	}

	//get config
	content, err := client.GetConfig(vo.ConfigParam{
		DataId: "curd-config",
		Group:  "nooocode",
	})
	if err != nil {
		panic(err)
	}
	err = yaml.UnmarshalYML([]byte(content), DefaultConfig)
	if err != nil {
		panic(err)
	}
}

type Config struct {
	Mysql            string   `yaml:"mysql"`
	Debug            bool     `yaml:"debug"`
	PlatformTenantID string   `yaml:"platformTenantID"`
	EnableTenant     bool     `yaml:"enableTenant"`
	BasicForm        []string `yaml:"basicForm"`
	BasicPage        []string `yaml:"basicPage"`
}
