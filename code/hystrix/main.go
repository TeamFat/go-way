package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/afex/hystrix-go/hystrix"
	"github.com/astaxie/beego/logs"
)

func main() {
	//https://studygolang.com/articles/11875
	logs.Info("start")
	hystrix.Go("get_baidu", func() error {
		// talk to other services
		_, err := http.Get("https://www.baidu1.com/")
		if err != nil {
			fmt.Println("get error")
			return err
		}
		logs.Info("ok")
		return nil
	}, func(err error) error {
		logs.Info("get an error, handle it")
		return nil
	})

	time.Sleep(2 * time.Second)
}
