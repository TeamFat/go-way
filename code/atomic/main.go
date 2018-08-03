package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

var (
	count int32
	wg    sync.WaitGroup
)

func main() {
	go func() {
		log.Println(http.ListenAndServe("localhost:8081", nil))
	}()
	wg.Add(2)
	go incCount(1)
	go incCount(2)
	wg.Wait()
	fmt.Println("dfdadf", count)
	time.Sleep(time.Minute)

}

func incCount(ii int32) {
	defer wg.Done()
	for i := 0; i < 2; i++ {
		runtime.Gosched()
		atomic.AddInt32(&count, 1)
	}
}
