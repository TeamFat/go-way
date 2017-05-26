package main

import "fmt"

func main() {
	x := []int{1, 2, 3}
	y := []int{4, 5, 6}
	//注意下面这两个区别
	fmt.Println(append(x, 4, 5, 6))
	fmt.Println(append(x, y...))
}

// 输出结果：

// [1 2 3 4 5 6]
// [1 2 3 4 5 6]

// append的用法有两种：
// slice = append(slice, elem1, elem2)
// slice = append(slice, anotherSlice…)

// 第一种用法中，第一个参数为slice,后面可以添加多个参数。
// 如果是将两个slice拼接在一起，则需要使用第二种用法，在第二个slice的名称后面加三个点，而且这时候append只支持两个参数，不支持任意个数的参数。
