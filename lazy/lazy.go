package lazy

import (
// "fmt"
)

// 添加路由
func AddRouter(path string, c ControllerInterface, a string) *App {
	LazyApp.Router.Add(path, c, a)

	return LazyApp
}

func Run() {
	LazyApp.Run()
}
