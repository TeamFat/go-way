package main

import (
	"flag"
	"fmt"
)

func main() {
	// 待使用的变量
	var id int
	var name string
	var male bool

	//设置flag参数 (变量指针，参数名，默认值，帮助信息)
	flag.IntVar(&id, "id", 123, "help msg for id")
	flag.StringVar(&name, "name", "default name", "help msg for name")
	flag.BoolVar(&male, "male", false, "help msg for male")

	// 解析
	flag.Parse()

	// flag参数
	fmt.Printf("id = %d\n", id)
	fmt.Printf("name = %s\n", name)
	fmt.Printf("male = %t\n", male)

	//输入& 输出
	//	go run main.go -id=1234 -name="chenjie" -male=true
	//	id = 1234
	//	name = chenjie
	//	male = true

}
