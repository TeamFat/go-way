package main

import (
	"fmt"
	"sync"
	"time"
)

var m sync.Mutex

func main() {
	//带new的例子 http://www.baiyuxiong.com/?cat=22&paged=10
	//new make的区别http://studygolang.com/articles/3496
	//m = new(sync.Mutex)

	go lockPrint(1)
	lockPrint(2)

	time.Sleep(time.Second)

	fmt.Printf("%s\n", "exit!")
}

func lockPrint(i int) {
	println(i, "lock start")

	m.Lock()

	println(i, "in lock")
	time.Sleep(3 * time.Second)

	m.Unlock()
	println(i, "unlock")
}
