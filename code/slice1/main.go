package main

import (
	"fmt"
)

func main() {
	slice := make([]int, 3, 5)
	appendTest(slice)
	fmt.Println(slice)
	//slice函数传参是传值不是传引用
	//[0 0 0]
}

func appendTest(slice []int) {
	slice = append(slice, 5)
}
