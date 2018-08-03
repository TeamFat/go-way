package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var (
	count int32
	wg    sync.WaitGroup
)

func main() {
	wg.Add(2)
	go incCount(1)
	go incCount(2)
	wg.Wait()
	fmt.Println("dfdadf", count)
}

func incCount(ii int32) {
	defer wg.Done()
	for i := 0; i < 2; i++ {
		runtime.Gosched()
		atomic.AddInt32(&count, 1)
	}
}
