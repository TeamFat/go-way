package main

import (
	"fmt"
	"time"
)

type Job int64

func process(job Job) {
	fmt.Println(job)
}

func worker(jobChan <-chan Job, cancelChan <-chan struct{}) {
	for {
		select {
		case _, ok := <-cancelChan:
			fmt.Println(ok)
			return

		case job := <-jobChan:
			process(job)
		}
	}
}
func main() {
	cancelChan := make(chan struct{})
	jobChan := make(chan Job, 10)
	for i := 0; i < 10; i++ {
		jobChan <- Job(i)
	}
	// start the goroutine passing it the cancel channel
	go worker(jobChan, cancelChan)

	// to cancel the worker, close the cancel channel
	close(cancelChan)
	time.Sleep(5 * time.Second)
	//这个例子有公平问题，https://www.opsdash.com/blog/job-queues-in-go.html
}
