package dao

import (
	"studysystem_micro/internal/service/user/model"
	"studysystem_micro/pkg/init/sql"
)

// 创建用户
func Create_user(a *model.User) {
	db := sql.GetDB()
	db.Create(a)
}

// 查询用户

// 根据email查询用户,通常在登录时用
func Search_user(email string) *model.User {
	db := sql.GetDB()
	user := new(model.User)
	db.Where("email = ?", email).Find(user)
	return user
}

// 根据id查询用户
func Search_user_by_id(id int64) *model.User {
	db := sql.GetDB()
	user := new(model.User)
	db.Where("id = ?", id).Find(user)
	return user
}

// 修改用户
// action:
// 1-->nickname,2-->avatar,3-->individualresume
func Update_user(id int64, action int32, data string) {
	db := sql.GetDB()
	u := Search_user_by_id(id)
	switch action {
	case 1:
		u.NickName = data
	case 2:
		u.Avatar = data
	case 3:
		u.IndividualResume = data
	}
	db.Save(u)
}
