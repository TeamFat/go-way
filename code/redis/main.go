package main

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

func main() {
	c, err := redis.Dial("tcp", "ap2.jd.local:5360")
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	}
	if _, err := c.Do("AUTH", "/redis/cluster/24:1417694442520"); err != nil {
		c.Close()
		return
	}
	defer c.Close()
	c.Send("SET", "foo", "bar")
	c.Send("GET", "foo")
	c.Send("GET", "foo1")
	c.Flush()
	c.Receive()                         // reply from SET
	v, err := redis.String(c.Receive()) // reply from GET
	fmt.Println(v, err)
	fmt.Println(len(v))
	v1, err1 := redis.String(c.Receive()) // reply from GET
	fmt.Println(v1 == "")
	fmt.Println(err1)
	//var abc []string
	abc := []string{v, v1}
	fmt.Println(len(abc))
	var baseArr [15]string
	baseArr[0] = "abc"
	baseArr[14] = "abc"
	fmt.Println(len(baseArr))
	fmt.Println(baseArr[0], baseArr[1], baseArr[2], baseArr[14])
	fmt.Println(baseArr[1])
}
