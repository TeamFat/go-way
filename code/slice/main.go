package main

import (
	"fmt"
)

//删除函数
func remove(s []string, i int) []string {
	return append(s[:i], s[i+1:]...)
}

func main() {
	s := []string{"a", "b", "c"}
	fmt.Println(s)
	s = remove(s, 1)
	fmt.Println(s)
}
