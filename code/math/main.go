package main

import (
	"fmt"
	"math"
)

func main() {
	//2为底值8的对数
	fNum := math.Log2(8)
	fmt.Println(fNum)
	//取浮点的整数和小数部分
	f, frac := math.Modf(fNum)
	//判断对数是否为整数
	if frac == 0 {
		fmt.Println(f)
	}
}
