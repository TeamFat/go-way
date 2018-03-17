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

	var d float64 = 1129.6
	var e int64 = int64(d * 100)
	fmt.Println(e)               //output:112959
	fmt.Println(Round(d*100, 0)) //output:112960
	//http://blog.csdn.net/sjy8207380/article/details/79013827
	//https://imhanjm.com/2017/08/27/go%E5%A6%82%E4%BD%95%E7%B2%BE%E7%A1%AE%E8%AE%A1%E7%AE%97%E5%B0%8F%E6%95%B0-decimal%E7%A0%94%E7%A9%B6/
}

func Round(f float64, n int) float64 {
	n10 := math.Pow10(n)
	return math.Trunc((f+0.5/n10)*n10) / n10
}
