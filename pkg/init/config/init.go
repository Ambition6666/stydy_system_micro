package config

import (
	"fmt"
	"log"
	"reflect"

	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

const (
	nacos_Host = "192.168.1.75"
	nacos_Port = 8848
)

// 每个服务初始化配置
func Init_config(data_id string, group string, value any) {
	sc := []constant.ServerConfig{{
		IpAddr: nacos_Host,
		Port:   nacos_Port,
	}}

	cc := constant.ClientConfig{
		NamespaceId:         "232e51d6-1528-43b4-ab13-aa69039c7886", // 如果需要支持多namespace，我们可以场景多个client,它们有不同的NamespaceId。当namespace是public时，此处填空字符串。
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "log",
		CacheDir:            "cache",
		LogLevel:            "debug",
	}

	configClient, err := clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": sc,
		"clientConfig":  cc,
	})
	if err != nil {
		fmt.Println(err.Error())
	}

	content, err := configClient.GetConfig(vo.ConfigParam{
		DataId: data_id,
		Group:  group,
	})
	SetConfig(content, value)

	if err != nil {
		fmt.Println(err.Error())
	}
	err = configClient.ListenConfig(vo.ConfigParam{
		DataId: data_id,
		Group:  group,
		OnChange: func(namespace, group, dataId, data string) {
			fmt.Println("配置文件发生了变化...")
			fmt.Println("group:" + group + ", dataId:" + dataId + ", data:" + data)
			SetConfig(data, value)
		},
	})
	if err != nil {
		fmt.Println(err.Error())
	}
}

func SetConfig(content any, value any) {
	tval := reflect.TypeOf(value)
	if tval.Kind() != reflect.Ptr {
		log.Fatal("非指针类型")
		return
	}
	rval := reflect.ValueOf(value)
	val := reflect.ValueOf(content)
	rval.Method(0).Call([]reflect.Value{val})
}
