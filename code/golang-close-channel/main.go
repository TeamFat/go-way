package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	jobs := make(chan int)
	timeout := make(chan bool)
	var wg sync.WaitGroup
	go func() {
		time.Sleep(time.Second * 3)
		timeout <- true
	}()
	go func() {
		for i := 0; ; i++ {
			select {
			case <-timeout:
				close(jobs)
				return
			default:
				jobs <- i
				fmt.Println("produce:", i)
			}
		}
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := range jobs {
			fmt.Println("consume:", i)
		}
	}()
	wg.Wait()
	//https://blog.tanteng.me/2017/11/golang-close-channel/
}
