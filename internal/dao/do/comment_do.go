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

/*
 * @Package do
 * @Author Quan Chen
 * @Date 2020/3/19
 * @Description
 *
 */
package do

import "github.com/opensourceai/go-api-service/internal/models"

/*评论*/
type CommentDO struct {
	models.Comment
	FromUser UserDo `json:"from_user" gorm:"foreignkey:FromUserId"`
	ToUser   UserDo `json:"to_user" gorm:"foreignkey:ToUserId"`
}

func (CommentDO) TableName() string {
	return "comment"
}
