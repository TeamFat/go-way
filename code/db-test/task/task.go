package task

import (
	"db-test/db"
	"fmt"

	"github.com/garyburd/redigo/redis"
)

func Run() {
	db.C.Do("SET", "hello", "world")
	s, _ := redis.String(db.C.Do("GET", "hello"))
	fmt.Printf("%#v\n", s)
}
