package main

import "fmt"
import "time"
import "sync/atomic"
import "runtime"

func main() {

	// 定义一个整数
	var ops uint64 = 0

	// 使用50个线程给ops累加数值
	for i := 0; i < 50; i++ {
		go func() {
			for {
				// 每次加1
				atomic.AddUint64(&ops, 1)

				// 这个函数用于时间片切换
				//可以理解为高级版的time.Sleep()
				//避免前面的for循环将CPU时间片都卡在一个线程里，使得其它线程没有执行机会
				runtime.Gosched()
			}
		}()
	}

	//停一秒，上面50个线程有1秒的执行时间
	time.Sleep(time.Second)

	// 获取结果
	opsFinal := atomic.LoadUint64(&ops)
	fmt.Println("ops:", opsFinal)
}
