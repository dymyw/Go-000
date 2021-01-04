package service

import (
	"Week02/biz"
	"fmt"
)

// service -> biz -> dao

func BatchGetUser() {
	user, err := biz.BatchGetUser()
	if err != nil {
		fmt.Printf("%+v\n", err)
		return
	}

	fmt.Println(user)
}
