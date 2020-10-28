package main

import (
	"fmt"
	"gitflow/login"
)

func main() {
	isSuccess := login.Login()
	if isSuccess {
		fmt.Println("login success!")
	} else {
		fmt.Println("login failed!")
	}
}
