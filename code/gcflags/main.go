package main

import (
	"fmt"
)

func main() {
	s := "hello"
	fmt.Println(s)
}

//全局变量逃逸
// go run -gcflags '-m -l' main.go
// # command-line-arguments
// ./main.go:9: s escapes to heap
// ./main.go:9: main ... argument does not escape
// hello
