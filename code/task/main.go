package main

import (
	"os"
	"time"

	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/toolbox"
)

func init() {
	os.MkdirAll("./log/", 0777)
	//日志默认不输出调用的文件名和文件行号,如果你期望输出调用的文件名和文件行号,可以如下设置
	//开启传入参数 true,关闭传入参数 false,默认是关闭的.
	logs.EnableFuncCallDepth(true)
	//如果你的应用自己封装了调用 log 包,那么需要设置 SetLogFuncCallDepth,默认是 2,也就是直接调用的层级,如果你封装了多层,那么需要根据自己的需求进行调整.
	logs.SetLogFuncCallDepth(3)
	logs.SetLogger("file", `{"filename":"./log/test.log"}`)
	logs.SetLevel(logs.LevelDebug)
}

func main() {
	// NewTask(tname string, spec string, f TaskFunc) *Task

	// tname 任务名称
	// spec 定时任务格式，请参考下面的详细介绍
	// f 执行的函数 func() error
	tk1 := toolbox.NewTask("tk1", "0/1 * * * * *", func() error { logs.Info("test"); return nil })
	toolbox.AddTask("tk1", tk1)
	toolbox.StartTask()
	time.Sleep(time.Second * 10)
	defer toolbox.StopTask()
}
