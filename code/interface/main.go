package main

import (
	"fmt"
)

//1
type I interface {
	Get() int
	Set(int)
}

//2
type S struct {
	Age int
}

func (s S) Get() int {
	return s.Age
}

func (s *S) Set(age int) {
	s.Age = age
}

type SS struct {
	Age int
}

func (s SS) Get() int {
	return s.Age
}

func (s SS) Set(age int) {
	s.Age = age
}

//3
func f(i I) {
	i.Set(10)
	fmt.Println(i.Get())
}

type R struct{ i int }

func (p *R) Get() int  { return p.i }
func (p *R) Set(v int) { p.i = v }

func doSomething(v interface{}) {
	fmt.Println("doSomething called")
}

func printAll(vals []interface{}) {
	for _, val := range vals {
		fmt.Println(val)
	}
}

type Animal interface {
	Speak() string
}

type Dog struct {
}

func (d Dog) Speak() string {
	return "Woof!"
}

type Cat struct {
}

func (c *Cat) Speak() string {
	return "Meow!"
}

type Llama struct {
}

func (l Llama) Speak() string {
	return "?????"
}

type JavaProgrammer struct {
}

func (j JavaProgrammer) Speak() string {
	return "Design patterns!"
}

func main() {
	//http://sanyuesha.com/2017/07/22/how-to-understand-go-interface/
	s := S{}
	f(&s) //4

	ss := SS{}
	f(&ss) //ponter
	f(ss)  //value

	var i I
	i = &s
	fmt.Println(i.Get())

	if _, ok := i.(*S); ok {
		fmt.Println("i store *S")
	}

	switch t := i.(type) {
	case *S:
		fmt.Println("i store *S", t)
	case *R:
		fmt.Println("i store *R", t)
	}

	doSomething(&s)

	names := []string{"stanley", "david", "oscar"}
	vals := make([]interface{}, len(names))
	for index, value := range names {
		vals[index] = value
	}
	// printAll(names)
	printAll(vals)

	//cannot use Cat literal (type Cat) as type Animal in array or slice literal:
	// Cat does not implement Animal (Speak method has pointer receiver)
	//
	// animals := []Animal{Dog{}, Cat{}, Llama{}, JavaProgrammer{}}
	// for _, animal := range animals {
	// 	fmt.Println(animal.Speak())
	// }
}
