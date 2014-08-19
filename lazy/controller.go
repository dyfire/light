package lazy

import (
	"errors"
	"fmt"
	"net/http"
	"reflect"
)

var (
	controller Controller
)

type Controller struct {
	// Handler map[string]func()
	Handler map[string]string
	Writer  http.ResponseWriter
	Request *http.Request
}

type ControllerInterface interface {
	Init()
	Post()
}

func (c *Controller) Post() {

}

func (c *Controller) Init(path, method string) {
	c.Handler[path] = method
}

func (c *Controller) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c.Writer = w
	c.Request = r

	c.call(r.URL.Path)
}

func (c *Controller) AddRouter(router string, f func()) {
	// fmt.Println(c)
	// c.Handler[router] = f
	// fmt.Println("add", c.Handler)
}

func NewController() *Controller {
	return &Controller{
		Handler: make(map[string]func()),
		Writer:  nil,
		Request: nil,
	}
}

// reflect call func
func (c *Controller) call(name string, params ...interface{}) (result []reflect.Value, err error) {
	f := c.Handler
	if _, ok := f[name]; ok {
		v := reflect.ValueOf(f[name])
		if v.Type().NumIn() != len(params) {
			err = errors.New("The number of params is not adapted")
			return
		}

		in := make([]reflect.Value, len(params))
		for k, v := range params {
			in[k] = reflect.ValueOf(v)
		}

		result = v.Call(in)
	}

	return
}
