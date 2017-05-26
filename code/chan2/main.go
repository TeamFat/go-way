package main

import (
	"fmt"
)

func main() {
	c := make(chan func())

	go func() {
		for i := 0; i < 10; i++ {
			c <- f
		}
		close(c)
	}()

	for f1 := range c {
		f1()
	}

}

func f() {
	fmt.Println(1)
}
