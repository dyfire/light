package router

import (
	"fmt"
	"light/controller"
	"light/lazy"
)

func init() {
	fmt.Println("start router")
	lazy.AddRouter("/user", &controller.UserController{}, "list")
	lazy.AddRouter("/post", &controller.PostController{}, "list")
}
