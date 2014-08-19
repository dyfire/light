package lazy

import (
	"fmt"
	// "net/http"
	// "path/filepath"
	// "strconv"
	// "log"
)

type App struct {
	Router *Router
}

func NewApp() *App {
	fmt.Println("init app")
	router := NewRouter()
	fmt.Println(router)
	return &App{
		Router: router,
	}
}

func (ap *App) Run() {
	// fmt.Println("server start")
	// err := http.ListenAndServe(":"+fmt.Sprintf("%d", HttpPort), ap.Handlers)

	// if err != nil {
	// 	log.Fatal(err)
	// }
}
