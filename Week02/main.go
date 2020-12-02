package main

import (
	"Week02/service"
	"fmt"
)

func main() {
	id := 1
	user, err := service.GetUserById(id)
	if err != nil {
		fmt.Printf("%+v", err)
	}

	fmt.Println(user)
}
