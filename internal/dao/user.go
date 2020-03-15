package dao

import "github.com/opensourceai/go-api-service/internal/models"

type UserDao interface {
	// 添加用户
	Add(user *models.User) error
	// 根据id软删除用户
	DeleteById(ids ...int) error
	// 编辑用户信息
	Edit(user *models.User) error
	// 获取该用户名的用户信息
	GetUserByUsername(username string) (err error, user models.User)
}