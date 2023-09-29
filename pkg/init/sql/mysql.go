package sql

import (
	"fmt"
	"reflect"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func GetDB() *gorm.DB {
	return db
}
func InitMysql(val any) {
	t := reflect.TypeOf(val)
	if t.Kind() != reflect.Struct {
		fmt.Println("非结构体类型")
		return
	}
	v := reflect.ValueOf(val)
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", v.Field(0).String(), v.Field(1).String(), v.Field(2).String(), v.Field(3).String(), v.Field(4).String())
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		fmt.Println(err)
	}
}
