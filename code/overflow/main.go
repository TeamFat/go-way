package main

import "fmt"

func se() func() int {
	return func() int {
		var x int
		x++
		return x * x
	}
}

func main() {
	f := se()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
}

//go run -gcflags "-m -l" main.go
