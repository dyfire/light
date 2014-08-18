package router

import (
	// "light/controller"
	"fmt"
	"light/lazy"
)

func init() {
	fmt.Println("start router")
	lazy.AddRouter("/user", User)
	lazy.AddRouter("/login", Login)
}

func User() {
	fmt.Println("hello user")
}

func Login() {
	fmt.Println("login")
}
