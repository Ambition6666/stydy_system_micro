package sql

import (
	"fmt"
	"reflect"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var pgDB *gorm.DB

func Getpgsql() *gorm.DB {
	return pgDB
}
func InitPgsql(val any) {
	t := reflect.TypeOf(val)
	if t.Kind() != reflect.Struct {
		fmt.Println("非结构体类型")
		return
	}
	v := reflect.ValueOf(val)
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", v.Field(0).String(), v.Field(1).String(), v.Field(2).String(), v.Field(3).String(), v.Field(4).String())
	var err error
	pgDB, err = gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
}
