package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {

	//返回的是指针需要特别注意
	// Int(name string, value int, usage string) *int
	// String(name string, value string, usage string) *string
	// Bool(name string, value bool, usage string) *bool
	var (
		id   = flag.Int("id", 123, "help msg for id")
		name = flag.String("name", "default name", "help msg for name")
		male = flag.Bool("male", false, "help msg for male")
	)

	// 解析
	flag.Parse()

	// flag参数
	fmt.Printf("id = %d\n", *id)
	fmt.Printf("name = %s\n", *name)
	fmt.Printf("male = %t\n", *male)
	os.Exit(0)

	//输入& 输出
	//	go run main.go -id=1234 -name="chenjie" -male=true
	//	id = 1234
	//	name = chenjie
	//	male = true

}
