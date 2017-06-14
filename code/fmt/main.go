package main

import "fmt"

func main() {
	var i float64 = -2.30
	fmt.Printf("%g", i)
	fmt.Printf("%G", i)
	fmt.Printf("%t", true)
	fmt.Printf("%%")
	fmt.Printf("%q", "Go语言中文网")
	//-2.3-2.3true%"Go语言中文网"
	//https://github.com/polaris1119/The-Golang-Standard-Library-by-Example/blob/master/chapter01/01.3.md
}
