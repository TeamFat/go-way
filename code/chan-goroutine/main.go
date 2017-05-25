//本文通过代码讲解如何实现一个线程池。代码及注释如下：
package main

import "fmt"
import "time"

//这个是工作线程，处理具体的业务逻辑，将jobs中的任务取出，处理后将处理结果放置在results中。
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "processing job", j)
		time.Sleep(time.Second)
		results <- j * 2
	}
}

func main() {

	//两个channel，一个用来放置工作项，一个用来存放处理结果。
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	// 开启三个线程，也就是说线程池中只有3个线程，实际情况下，我们可以根据需要动态增加或减少线程。
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// 添加9个任务后关闭Channel
	// channel to indicate that's all the work we have.
	for j := 1; j <= 9; j++ {
		jobs <- j
	}
	close(jobs)

	//获取所有的处理结果
	for a := 1; a <= 9; a++ {
		<-results
	}
}

// 输出结果：
// worker 1 processing job 1
// worker 2 processing job 2
// worker 3 processing job 3
// worker 1 processing job 4
// worker 3 processing job 5
// worker 2 processing job 6
// worker 1 processing job 7
// worker 2 processing job 8
// worker 3 processing job 9

// 从中可以看出，多个线程轮流处理了9个任务。
// 通过这个例子，我们可以学习到：
// 1、GO中多线程应用开发非常简单。
// 2、Channel是不同线程间数据交互的利器。 上面的例子中，主线程向jobs中写数据，三个工作线程同时从一个Channel中取数据。
