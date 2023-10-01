package rpc

import (
	"bytes"

	"github.com/spf13/viper"
)

type ServerConfig struct {
	Consul CConsul `yaml:"consul"`
	User CUser `yaml:"user"`
}

func (a *ServerConfig) BindData(val string) {
	var runtime_viper = viper.New()
	runtime_viper.SetConfigType("yaml")
	runtime_viper.ReadConfig(bytes.NewBuffer([]byte(val)))
	runtime_viper.Unmarshal(a)
}

var C *ServerConfig

type CConsul struct{
	Host string `yaml:"host"`
	port string `yaml:"port"`
}

type CUser struct{
	Name string `yaml:"name"`
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}