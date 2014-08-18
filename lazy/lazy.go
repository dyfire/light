package lazy

import (
// "fmt"
)

// 添加路由
func AddRouter(path string, f func()) *App {
	LazyApp.Handlers.AddRouter(path, f)
	return LazyApp
}

func Run() {
	LazyApp.Run()
}
