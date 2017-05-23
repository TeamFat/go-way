package main

import (
	"github.com/astaxie/beego/logs"
)

func main() {
	//an official log.Logger
	l := logs.GetLogger()
	l.Println("this is a message of http")
	//an official log.Logger with prefix ORM
	logs.GetLogger("ORM").Println("this is a message of orm")

	logs.Debug("my book is bought in the year of ", 2016)
	logs.Info("this %s cat is %v years old", "yellow", 3)
	logs.Warn("json is a type of kv like", map[string]int{"key": 2016})
	logs.Error(1024, "is a very", "good game")
	logs.Critical("oh,crash")
	log := logs.NewLogger()
	log.SetLogger(logs.AdapterConsole)
	log.Debug("this is a debug message")

	// 2017/05/24 01:07:34 [ORM] this is a message of orm
	// 2017/05/24 01:07:34 [D] my book is bought in the year of  2016
	// 2017/05/24 01:07:34 [I] this yellow cat is 3 years old
	// 2017/05/24 01:07:34 [W] json is a type of kv like map[key:2016]
	// 2017/05/24 01:07:34 [E] 1024 is a very good game
	// 2017/05/24 01:07:34 [C] oh,crash
	// 2017/05/24 01:07:34 [D] this is a debug message
}
