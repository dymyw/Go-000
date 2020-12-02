package main

import (
	"Week02/service"
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
)

func main() {
	id := 1
	user, err := service.GetUserById(id)

	// 404
	if errors.Is(err, sql.ErrNoRows) {
		fmt.Println("Not Found")
		return
	}

	if err != nil {
		fmt.Printf("%+v", err)
		return
	}

	fmt.Println(user)
}
