package lazy

import (
	"fmt"
	"reflect"
)

type Router struct {
	Control map[string]ControllerInterface
}

func NewRouter() *Router {
	fmt.Println("router start")
	return &Router{
		Control: make(map[string]ControllerInterface),
	}
}

func (r *Router) Add(path string, c ControllerInterface, action string) {
	reflectVal := reflect.ValueOf(c)

	// Indirect returns the value that v points to.
	// If v is a nil pointer, Indirect returns a zero Value.
	// If v is not a pointer, Indirect returns v.
	t := reflect.Indirect(reflectVal).Type()

	controllerName := string(t.Name())
	c.Init(path, action)
	// c.Handler[path] = action
	r.Control[controllerName] = c
	// c.Handler
	//

	fmt.Println(r)
}
