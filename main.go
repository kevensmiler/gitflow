package main

import (
	"fmt"
	"gitflow/register"
)

func main() {
	isSuccess := register.Register()
	if isSuccess {
		fmt.Println("register success!")
	} else {
		fmt.Println("register failed!")
	}
}
