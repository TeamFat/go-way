package main

import (
	"flag"
	"os"

	"github.com/astaxie/beego/logs"
)

var (
	topic   = flag.String("topic", "", "nsq topic.")
	channel = flag.String("channel", "Alarm", "nsq channel.")
	lookups = flag.String("lookups", "", "nsqlookup addrs: ip:port,ip:port.")
)

func init() {
	flag.Parse()

	os.MkdirAll("./log/", 0777)
	//日志默认不输出调用的文件名和文件行号,如果你期望输出调用的文件名和文件行号,可以如下设置
	//开启传入参数 true,关闭传入参数 false,默认是关闭的.
	logs.EnableFuncCallDepth(true)
	//如果你的应用自己封装了调用 log 包,那么需要设置 SetLogFuncCallDepth,默认是 2,也就是直接调用的层级,如果你封装了多层,那么需要根据自己的需求进行调整.
	logs.SetLogFuncCallDepth(3)
	logs.SetLogger("file", `{"filename":"./log/alarm.log"}`)
	logs.SetLevel(logs.LevelDebug)
}

func printParam() {

	logs.Info("using param:")
	logs.Info("lookups:", *lookups)
	logs.Info("topic:", *topic)
	logs.Info("channel:", *channel)
}

func main() {
	printParam()
	logs.Info("Alarm started.")
	/*
		2017/05/24 13:46:37 [I] [main.go:28] using param:
		2017/05/24 13:46:37 [I] [main.go:29] lookups:
		2017/05/24 13:46:37 [I] [main.go:30] topic:
		2017/05/24 13:46:37 [I] [main.go:31] channel: Alarm
		2017/05/24 13:46:37 [I] [main.go:36] Alarm started.
	*/
}
