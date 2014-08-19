package main

import (
	"fmt"
	// "runtime"
	// "path/filepath"
	"reflect"
	// "strings"
	// "strconv"
)

type User struct {
	name string
	age  int
}

type Notifier interface {
	Notify() error
}

func (u *User) Notify() error {
	return nil
}

func main() {
	u := &User{}

	v := reflect.ValueOf(u)
	// k := reflect.TypeOf(u)
	e := v.Interface()
	// fmt.Println(string(v.String()))
	fmt.Println(e)
}
