package service

import (
	"Week02/dao"
	"Week02/model"
	"github.com/pkg/errors"
)

func GetUserById(id int) (*model.User, error) {
	user, err := dao.GetUserById(id)
	if err != nil {
		return nil, errors.WithMessage(err, "service GetUserById failed!")
	}

	return user, nil
}
