package lazy

import (
	"fmt"
)

var (
	LazyApp  *App
	HttpAddr string
	HttpPort int
)

func ParseConfig() {

}

func init() {
	fmt.Println("lazy config")
	LazyApp = NewApp()

	HttpPort = 8888
}
