package biz

import (
	"Week02/code"
	"Week02/dao"
	"errors"
)

var UserNotFound = errors.New("user not found")

func BatchGetUser() ([]dao.User, error) {
	users, err := dao.BatchGetUser()
	if err != nil {
		// 在业务层通过 errors.Is(err, code.NotFound) 进行判断
		if errors.Is(err, code.NotFound) {
			//
			return nil, UserNotFound
		}
	}

	return users, nil
}
