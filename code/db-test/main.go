package main

import (
	"time"

	"github.com/TeamFat/go-way/code/db-test/db"
	"github.com/TeamFat/go-way/code/db-test/task"
)

func main() {
	db.ConnRedis()
	go task.Run()
	//go task.RunT()
	//不支持并发，需要改成redis连接池
	time.Sleep(time.Second * 3)
	db.StopRedis()
}
