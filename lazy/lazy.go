package lazy

import (
// "fmt"
)

func AddRouter(path string, f func()) *App {
	LazyApp.Handlers.AddRouter(path, f)
	return LazyApp
}

func Run() {
	LazyApp.Run()
}
