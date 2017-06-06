package task

import (
	"fmt"

	"github.com/TeamFat/go-way/code/db-test/db"
	"github.com/garyburd/redigo/redis"
)

func Run() {
	db.C.Do("SET", "hello", "world")
	s, _ := redis.String(db.C.Do("GET", "hello"))
	fmt.Printf("%#v\n", s)
}

func RunT() {
	db.C.Do("SET", "helloT", "world")
	s, _ := redis.String(db.C.Do("GET", "helloT"))
	fmt.Printf("%#v\n", s)
}
