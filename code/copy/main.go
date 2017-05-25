// Go内建函数copy:
// func copy(dst, src []Type) int
// 用于将源slice的数据（第二个参数），复制到目标slice（第一个参数）。
// 返回值为拷贝了年数据个数，是len(dst)和len(src)中的最小值。
package main

import (
	"fmt"
)

func main() {
	var a = []int{0, 1, 2, 3, 4, 5, 6, 7}
	var s = make([]int, 6)

	//源长度为8，目标为6，只会复制前6个
	n1 := copy(s, a)
	fmt.Println("s - ", s)
	fmt.Println("n1 - ", n1)

	//源长为7，目标为6，复制索引1到6
	n2 := copy(s, a[1:])
	fmt.Println("s - ", s)
	fmt.Println("n2 - ", n2)

	//源长为8-5=3，只会复制3个值，目标中的后三个值不会变
	n3 := copy(s, a[5:])
	fmt.Println("s - ", s)
	fmt.Println("n3 - ", n3)

	//将源中的索引5,6,7复制到目标中的索引3,4,5
	n4 := copy(s[3:], a[5:])
	fmt.Println("s - ", s)
	fmt.Println("n4 - ", n4)
}

// s -  [0 1 2 3 4 5]
// n1 -  6
// s -  [1 2 3 4 5 6]
// n2 -  6
// s -  [5 6 7 4 5 6]
// n3 -  3
// s -  [5 6 7 5 6 7]
// n4 -  3
