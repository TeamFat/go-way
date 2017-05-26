package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	filename := "test.log"
	data := []byte("hello")
	ioutil.WriteFile(filename, data, 0644)

	res, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(res))
}
