package main

import (
	"time"

	"github.com/astaxie/beego/logs"
)

func main() {
	jobResult := make(chan bool)
	go func() {
		time.Sleep(time.Second * 2)

		jobResult <- true
		//没机会读
		logs.Info("222222")

	}()
	t := time.NewTimer(1 * time.Second)
	select {
	case <-t.C:
		logs.Info("ccccc")
	case ret := <-jobResult:
		logs.Info(ret)
	}
	time.Sleep(time.Second * 10)
}
