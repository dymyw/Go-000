package dao

import (
	"Week02/model"
	"database/sql"
	"github.com/pkg/errors"
)

func GetUserById(id int) (*model.User, error) {
	// 模拟数据
	user := &model.User{}
	err := sql.ErrNoRows
	if err != nil {
		return nil, errors.Wrap(err, "dao GetUserById failed!")
	}

	return user, nil
}
