package main

import "fmt"

type RefIntArray2 *[2]int

func NewRefIntArray2() RefIntArray2 {
	return RefIntArray2(new([2]int))
}

func main() {
	refArr2 := NewRefIntArray2()
	fmt.Println(refArr2)
	modifyRefArr2(refArr2)
	fmt.Println(refArr2)
}

func modifyRefArr2(arr RefIntArray2) {
	arr[0] = 1
}

// 函数参数传值, 闭包传引用!
// slice 含 values/count/capacity 等信息, 是按值传递
// 按值传递的 slice 只能修改values指向的数据, 其他都不能修改
// slice 是结构体和指针的混合体
// 引用类型和传引用是两个概念
//http://studygolang.com/articles/4812
// &[0 0]
// &[1 0]
