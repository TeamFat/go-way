package main

import (
	"fmt"
)

func main() {
	s := "hello"
	fmt.Println(s)
}

//http://studygolang.com/articles/10026
//全局变量逃逸
// go run -gcflags '-m -l' main.go
// # command-line-arguments
// ./main.go:9: s escapes to heap
// ./main.go:9: main ... argument does not escape
// hello
