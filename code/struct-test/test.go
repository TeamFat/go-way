package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	m := make(map[int]*person)
	m[1] = &person{
		name: "chenjie",
		age:  29,
	}
	fmt.Println(m)
	m[1] = &person{
		name: "chenjie1",
		age:  28,
	}
	fmt.Println(m)
	fmt.Println(m[1].name)
	fmt.Println(m[1])
	m[1].name = "aa"
	m1 := make(map[int]person)
	m1[1] = person{
		name: "chenjie",
		age:  29,
	}
	fmt.Println(&m1[1]) //错误，不能取map的值地址，map值不能寻址
	fmt.Println(m1[1].name)
	m1[1].name = "bb" //除非值是指针类型，否则因为不能寻址所以不能对值属性赋值
}
