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
	Notify()
}

func (u *User) Notify() (s string) {
	fmt.Println("just do it")
	return "ok"
}

func (u *User) Writer() {

}

type D struct {
	F map[string]func()
}

func main() {
	u := &User{}
	// t := User{}
	v := reflect.ValueOf(u)

	e := v.MethodByName("Notify").Call(nil)
	// l := v.NumMethod()

	fmt.Println(e)

	a := reflect.TypeOf(u)
	e1 := a.Method(0)
	fmt.Println(e1.Func)

	d := &D{}
	d.F["f"] = e1.Func.Call(nil)
	fmt.Println(d)

}
