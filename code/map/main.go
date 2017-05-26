package main

import "fmt"

func main() {

	//使用make申请一个map，键为string类型，值为int类型
	m := make(map[string]int)

	//设置值
	m["k1"] = 7
	m["k2"] = 13

	//取指定键的值
	v1 := m["k1"]
	fmt.Println("v1: ", v1)

	//取长度
	fmt.Println("len:", len(m))

	//遍历
	for key, value := range m {
		fmt.Println(key, ":", value)
	}
	//删除
	delete(m, "k2")
	fmt.Println("map:", m)

	//初始化时直接指定值
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)
}
