package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	name string
	age  int
}

func Method(in interface{}) (ok bool) {
	v := reflect.ValueOf(in)
	if v.Kind() == reflect.Slice {
		ok = true
	} else {

	}
	num := v.Len()
	for i := 0; i < num; i++ {
		fmt.Println(v.Index(i).Interface())
	}
	return ok
}

func main() {
	s := []int{1, 3, 5, 7, 9}
	b := []float64{1.2, 3.4, 5.6, 7.8}
	Method(s)
	Method(b)
	fmt.Printf("%d\n", reflect.Bool) //output: 1
	fmt.Println(reflect.Bool)        //output: bool
}
