package main

import (
	"fmt"
)

type Sample struct {
	Name string
}

func SetName(s *Sample) {
	s.Name = "foo"
	fmt.Println("this is inside SetName")
	fmt.Println(s.Name)
}
func main() {
	s := &Sample{}
	s.Name = "bar"
	fmt.Println(s.Name)
	SetName(s)
	fmt.Println("this is outside SetName")
	fmt.Println(s.Name)
}
