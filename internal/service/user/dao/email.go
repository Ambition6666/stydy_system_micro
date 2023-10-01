package dao

import (
	"context"
	"studysystem_micro/pkg/init/sql"
	"time"
)

func SetAuthCode(em string, auth_code string) error {
	rdb := sql.GetRedis()
	return rdb.Set(context.Background(), "auth"+em, auth_code, 300*time.Second).Err()
}

// 获取验证码
func GetAuthCode(em string) (string, error) {
	rdb := sql.GetRedis()
	return rdb.Get(context.Background(), "auth"+em).Result()
}
