// Go提供了len内置函数，用于获取变量的长度：
// func len(v Type) int
// 根据变量v类型的不同，len返回值的意义也不同：
// 数组:：v中的元素的个数.
// 数组指针： *v中元素的个数
// Slice或map: v中的元素的个数; 如果v==nil, len(v)为0.
// 字符串: v中的字节数.
// Channel: the number of elements queued (unread) in the channel buffer;channel缓冲区中未读的元素个数。
// 如果v==nil, len(v)为0 .

// 在某些地方，cap函数和len很类似，所以我们在之前讲cap函数的示例代码中添加一些len的代码来看一下。
package main

import (
	"fmt"
)

func main() {
	arr := []int{1, 2, 3}
	fmt.Println("cap(arr) : ", cap(arr)) //3
	fmt.Println("len(arr) : ", len(arr)) //3

	fmt.Println()

	//长度和容量都为5
	slice1 := make([]string, 5)
	//长度为3，容量为5
	slice2 := make([]int, 3, 5)
	fmt.Println("cap(slice1) : ", cap(slice1)) //5
	fmt.Println("cap(slice2) : ", cap(slice2)) //5
	fmt.Println("len(slice1) : ", len(slice1)) //5
	fmt.Println("len(slice2) : ", len(slice2)) //3

	fmt.Println()

	c1 := make(chan int)
	c2 := make(chan int, 2)
	fmt.Println("cap(c1) : ", cap(c1)) //0
	fmt.Println("len(c1) : ", len(c1)) //0

	fmt.Println()

	fmt.Println("cap(c2) : ", cap(c2)) //2
	fmt.Println("len(c2) : ", len(c2)) //0

	fmt.Println()

	//c2中添加一个数
	c2 <- 1
	fmt.Println("cap(c2) : ", cap(c2)) //2
	fmt.Println("len(c2) : ", len(c2)) //1

	fmt.Println()

	s := "hello, world."
	fmt.Println("len(s) : ", len(s)) //13
	//cap不用用于string，下面的代码会报错：invalid argument s (type string) for cap
	//fmt.Println("cap(s) : ", cap(s))
}
