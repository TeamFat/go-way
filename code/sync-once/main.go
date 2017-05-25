package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	o := &sync.Once{}
	go do(o)
	go do(o)
	time.Sleep(time.Second * 2)
}

func do(o *sync.Once) {
	fmt.Println("Start do")
	o.Do(func() {
		fmt.Println("Doing something...")
	})
	fmt.Println("Do end")
}

// Start do
// Doing something...
// Do end
// Start do
// Do end

// 看go once的源码实现
// type Once struct {
//     m    Mutex
//     done uint32
// }

// func (o *Once) Do(f func()) {
//     if atomic.LoadUint32(&o.done) == 1 {
//         return
//     }
//     // Slow-path.
//     o.m.Lock()
//     defer o.m.Unlock()
//     if o.done == 0 {
//         defer atomic.StoreUint32(&o.done, 1)
//         f()
//     }
// }
// 核心思想是使用原子计数记录被执行的次数。使用Mutex Lock Unlock锁定被执行函数，防止被重复执行。
