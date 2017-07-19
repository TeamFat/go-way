package main

import (
	"fmt"
	"reflect"
)

type Foo struct {
	X string
	Y int
}

func main() {
	var i int = 123
	var f float32 = 1.23
	var l []string = []string{"a", "b", "c"}

	fmt.Println(reflect.TypeOf(i)) //int
	fmt.Println(reflect.TypeOf(f)) //float32
	fmt.Println(reflect.TypeOf(l)) //[]string

	var foo Foo
	fmt.Println(reflect.TypeOf(foo)) //main.Foo

}

/**
type Foo struct {
	X string
	Y int
}

func (f Foo) do() {
	fmt.Printf("X is: %s, Y is: %d", f.X, f.Y)

}

func main() {
	var s string = "abc"
	fmt.Println(reflect.TypeOf(s).String()) //string
	fmt.Println(reflect.TypeOf(s).Name())   //string

	var f Foo
	typ := reflect.TypeOf(f)
	fmt.Println(typ.String()) //main.Foo
	fmt.Println(typ.Name())   //Foo ，返回结构体的名字

}
**/
//Kind方法Type和Value都有，它返回的是对象的基本类型，例如int,bool,slice等，而不是静态类型。
//http://studygolang.com/articles/1251
