package lazy

import (
	"fmt"
	"net/http"
	// "path/filepath"
	// "strconv"
	"log"
)

type App struct {
	Handlers *Controller
}

func NewApp() *App {
	fmt.Println("init app")
	handler := NewController()
	return &App{
		Handlers: handler,
	}
}

func (ap *App) Run() {
	fmt.Println("server start")
	err := http.ListenAndServe(":"+fmt.Sprintf("%d", HttpPort), ap.Handlers)

	if err != nil {
		log.Fatal(err)
	}
}
