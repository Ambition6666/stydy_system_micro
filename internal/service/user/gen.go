package user

import (
	"bytes"

	"github.com/spf13/viper"
)

type Config struct {
	Email  Cemail  `yaml:"email"`
	Consul CConsul `yaml:"consul"`
}

var C *Config

func (a *Config) BindData(val string) {
	var runtime_viper = viper.New()
	runtime_viper.SetConfigType("yaml")
	runtime_viper.ReadConfig(bytes.NewBuffer([]byte(val)))
	runtime_viper.Unmarshal(a)
}

type Cemail struct {
	Addr   string `yaml:"addr"`
	Host   string `yaml:"host"`
	From   string `yaml:"from"`
	Email  string `yaml:"email"`
	Auth   string `yaml:"auth"`
	Expire int64  `yaml:"expire"`
}

type CConsul struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}
