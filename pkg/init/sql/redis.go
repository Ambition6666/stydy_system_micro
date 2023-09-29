package sql

import (
	"fmt"
	"reflect"

	re "github.com/redis/go-redis/v9"
)

var rdb *re.Client

func GetRedis() *re.Client {
	return rdb
}

func InitRedis(val any) {
	t := reflect.TypeOf(val)
	if t.Kind() != reflect.Struct {
		fmt.Println("非结构体类型")
		return
	}
	v := reflect.ValueOf(val)
	rdb = re.NewClient(&re.Options{
		Addr:     v.Field(0).String() + ":" + v.Field(1).String(),
		Password: v.Field(2).String(),
		DB:       0, // use default DB
	})
}
