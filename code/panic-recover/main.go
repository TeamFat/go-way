package main

import (
	"fmt"
)

func main() {
	fmt.Println("Start main")
	sub()
	fmt.Println("End main")
}

func sub() {
	defer handler()

	fmt.Println("Before panic")
	panic("golang_everyday")
	fmt.Println("After panic")
}

func handler() {
	if err := recover(); err != nil {
		fmt.Println("recover msg: ", err)
	} else {
		fmt.Println("recover ok")
	}
}

// 执行结果：
// Start main
// Before panic
// recover msg: golang_everyday
// End main
// 可以看到，sub函数没有执行完，panic后执行了defer就返回到上层main函数了，但是main函数执行结束了。
// 这是因为recover阻止了异常的继续传播。他将panic限制在了一定的范围内。
