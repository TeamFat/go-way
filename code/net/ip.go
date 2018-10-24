package main

import (
	"fmt"
	"net"
)

func main() {
	//判断是否ip
	ip := net.ParseIP("test.com") == nil
	fmt.Println(ip)
}
