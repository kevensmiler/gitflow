package main

import (
	"fmt"
	"gitflow/login"
	"gitflow/register"
)

func main() {
	isRegister := register.Register()
	if isRegister {
		fmt.Println("register success!")
	} else {
		fmt.Println("register failed!")
	}

	isLogin := login.Login()
	if isLogin {
		fmt.Println("login success!")
	} else {
		fmt.Println("login failed!")
	}
}
