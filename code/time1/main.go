package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now())
	fmt.Println(time.Now().Add(-time.Minute * 10))
	t := time.Now().Unix()
	fmt.Println(t - t%60)
	//更多https://zengweigang.gitbooks.io/core-go/eBook/04.8.html
}
