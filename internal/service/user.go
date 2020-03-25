/*
 *    Copyright 2020 opensourceai
 *
 *    Licensed under the Apache License, Version 2.0 (the "License");
 *    you may not use this file except in compliance with the License.
 *    You may obtain a copy of the License at
 *
 *        http://www.apache.org/licenses/LICENSE-2.0
 *
 *    Unless required by applicable law or agreed to in writing, software
 *    distributed under the License is distributed on an "AS IS" BASIS,
 *    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *    See the License for the specific language governing permissions and
 *    limitations under the License.
 */

package service

import (
	"errors"
	"github.com/google/wire"
	"github.com/opensourceai/go-api-service/internal/dao"
	"github.com/opensourceai/go-api-service/internal/dao/mysql"
	"github.com/opensourceai/go-api-service/internal/models"
	"github.com/opensourceai/go-api-service/pkg/util"
)

type UserService interface {
	Register(user *models.User) error
	Login(user models.User) (*models.User, bool, error)
}

type userService struct {
	dao.UserDao
}

var ProviderUser = wire.NewSet(NewUserService, mysql.NewUserDao)

func NewUserService(dao2 dao.UserDao) (UserService, error) {
	return &userService{dao2}, nil

}
func (service userService) Register(user *models.User) error {
	// 加密密码
	user.Password = util.EncodeMD5(user.Password)
	return service.DaoAdd(user)
}

func (service userService) Login(user models.User) (*models.User, bool, error) {
	err, u := service.DaoGetUserByUsername(user.Username)
	if err != nil {
		return nil, false, errors.New("登录失败")
	}
	// 匹配密码
	md5Password := util.EncodeMD5(user.Password)
	if u.Password == md5Password {
		return &u, true, nil
	}
	return nil, false, errors.New("登录失败")
}
