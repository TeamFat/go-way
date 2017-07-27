package main

import "fmt"

func main() {
	p1 := new(int)
	fmt.Printf("p1 --> %#v \n ", p1)           //(*int)(0xc42000e250)
	fmt.Printf("p1 point to --> %#v \n ", *p1) //0
	var p2 *int
	i := 0
	p2 = &i
	fmt.Printf("p2 --> %#v \n ", p2)           //(*int)(0xc42000e278)
	fmt.Printf("p2 point to --> %#v \n ", *p2) //0
	//http://sanyuesha.com/2017/07/26/go-make-and-new/
}
