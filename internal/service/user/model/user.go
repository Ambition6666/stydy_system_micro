package model

import "gorm.io/gorm"

type User struct {
	gorm.Model       `json:"-"`
	ID               int64  `json:"uid" gorm:"primarykey"`
	NickName         string `json:"nickname"`
	PassWord         string `json:"-"`
	Email            string `json:"email"`
	Avatar           string `json:"avatar"`
	IndividualResume string `json:"individual_resume"`
	Role             int    `json:"role"`
}

func NewUser(id int64, email string, password string) *User {
	u := &User{
		ID:               id,
		NickName:         "默认昵称",
		Email:            email,
		PassWord:         password,
		Avatar:           "avatar.jpg",
		Role:             0,
		IndividualResume: "个人简历",
	}
	return u
}
