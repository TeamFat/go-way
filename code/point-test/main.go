package main

import "fmt"

type test struct {
	num  int
	name string
}

func main() {

	i := test{num: 1, name: "chenjie"}
	fmt.Println(&i.name)
	b := &i
	fmt.Println(&(*b).name)
}
