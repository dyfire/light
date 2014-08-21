package main

import (
	"fmt"
	// "runtime"
	// "path/filepath"
	"reflect"
	// "strings"
	// "strconv"
	"log"
	"net/http"
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
	F map[string]reflect.Value
}

func (m *D) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("server http")

}

func main() {
	u := &User{}
	// t := User{}
	v := reflect.ValueOf(u)

	e := v.MethodByName("Notify")
	fmt.Println(e)

	d := D{
		F: make(map[string]reflect.Value),
	}
	d.F["a"] = e
	fmt.Println(d)
	err := http.ListenAndServe(":9090", &d) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
