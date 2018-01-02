package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	fmt.Println(rand.Intn(100)) //产生0-100的随机整数

	fmt.Println(rand.Float64()) //产生0.0-1.0的随机浮点点
	appid := "bm_core_order_i33sp"
	s1 := rand.NewSource(int64(len(appid))) //用指定值创建一个随机数种子
	r1 := rand.New(s1)
	fmt.Println(r1.Intn(59))
	fmt.Println(time.Now().UnixNano())

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		x := rand.Intn(100)
		fmt.Println(x)
	}
}
