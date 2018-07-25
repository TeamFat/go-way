package main

import (
	"fmt"
	"time"

	"github.com/astaxie/beego/logs"

	"golang.org/x/time/rate"
)

func main() {
	//https://studygolang.com/articles/13254

	//用户配置的平均发送速率为r，则每隔1/r秒一个令牌被加入到桶中（每秒会有r个令牌放入桶中），桶中最多可以存放b个令牌。如果令牌到达时令牌桶已经满了，那么这个令牌会被丢弃；
	//https://juejin.im/entry/5adef8305188256715474b53

	//https://www.jianshu.com/p/4ce68a31a71d
	l := rate.NewLimiter(10, 30)
	for {
		r := l.ReserveN(time.Now(), 8)
		logs.Info(r.Delay())
		time.Sleep(r.Delay())
		fmt.Println(time.Now().Format("2016-01-02 15:04:05.000"))
	}
}
