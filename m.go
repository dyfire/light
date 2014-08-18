package main

import (
	"fmt"
	// "runtime"
	// "path/filepath"
	"strings"
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
	fmt.Println(strings.Join(":a", "b"))
}
