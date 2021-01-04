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
		// 		Is errors 1.13 Unwrap => root cause error 能找到根因
		// 		biz error 处理业务逻辑错误
		// 		sql 不依赖 db、cache
		if errors.Is(err, code.NotFound) {
			return nil, UserNotFound
		}
	}

	return users, nil
}
