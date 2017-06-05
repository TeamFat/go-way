package main

import (
	"db-test/db"
	"db-test/task"
)

func main() {
	db.ConnRedis()
	task.Run()
	db.StopRedis()
}
