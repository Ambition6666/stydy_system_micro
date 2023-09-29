package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
	msk = []byte("ztynb6666")
)

// -------------------------------------jwt生成token加密------------------------------------------------
type Claim struct {
	ID   int64
	Role int
	jwt.RegisteredClaims
} //创建用户登录标签

// 得到token
func GetToken(id int64, role int) (string, error) {
	a := Claim{
		id,
		role,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(30 * 24 * time.Hour)), //token有效时间
			Issuer:    "zty",                                                   //签发人
		},
	} //获取claim实例
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, a) //获取token
	return token.SignedString(msk)                        //返回加密串
}

// 解析token
func ParseToken(token string) (*jwt.Token, int64, int, error) {
	claim := &Claim{}
	t, err := jwt.ParseWithClaims(token, claim, func(t *jwt.Token) (interface{}, error) {
		return msk, nil
	}) //接收前端发来加密字段
	return t, claim.ID, claim.Role, err
}
