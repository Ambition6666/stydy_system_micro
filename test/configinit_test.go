package test

import (
	"bytes"
	"fmt"
	"testing"

	"studysystem_micro/pkg/init/config"

	"github.com/spf13/viper"
)

type Test1 struct {
	Test `yaml:"test"`
}

type Test struct {
	A int `yaml:"a"`
}

func (a *Test1) BindData(c string) {
	fmt.Println(c)
	var runtime_viper = viper.New()
	runtime_viper.SetConfigType("yaml")
	runtime_viper.ReadConfig(bytes.NewBuffer([]byte(c)))
	runtime_viper.Unmarshal(a)
}
func TestXxx(t *testing.T) {
	v := new(Test1)
	config.Init_config("test", "DEFAULT_GROUP", v)
	fmt.Println(v)
}
