package main

import "fmt"
import "time"

//这个函数会在新的goroutine中运行，执行结束时会给done channel中传入值true
//注意到中间有sleep了一秒
func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	//channel中传入值true
	done <- true
}

func main() {

	//创建一个channel，用于goroutine之间通知情况
	done := make(chan bool, 1)
	//开启goroutine，并把done channel传进去
	go worker(done)

	//如果done channel中一直没有数据，这里就会卡住，直到worker结束时传入值以后，这里才会继续执行
	<-done
}

// 运行结果：

// working…done

// 如果我们把main函数最后一句代码：<-done，去掉以后，不会输出任何数据。
//因为主线程已经执行完退出了，goroutine依赖于主现程，也会退出。用古话说是： 皮之不存，毛將焉附？ 覆巢之下，安有完卵？ 国之不存，何以家为？
