package dao

import (
	"Week02/model"
	"database/sql"
)

func GetUserById(id int) (model.User, error) {
	// 模拟数据
	var user model.User
	err := sql.ErrNoRows

	return user, err
}
