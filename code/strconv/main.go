package main

import (
	"fmt"
	"strconv"
)

func main() {
	//第三个参数限制了转换后数值的大小，会有截断操作，尽管返回值最后会强制int64()
	y, _ := strconv.ParseInt("86400", 10, 64) // base 10, up to 64 bits
	fmt.Println(y)
}
