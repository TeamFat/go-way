package main

import "fmt"

var (
	aa = make(map[int]*string)
)

func main() {

	a := []string{"a", "b", "c"}
	for i, aStrategy := range a {
		fmt.Println(i)
		fmt.Println(aStrategy)
		b := aStrategy
		fmt.Println(&b)
		aa[i] = &b
	}
	fmt.Println(aa)
}
