package main

import (
	"fmt"
	"time"

	"github.com/garyburd/redigo/redis"
)

var RPool *redis.Pool

func main() {

	if err := ConnRedis("11.11.11.1:6379", 3); nil != err {
		fmt.Println(err)
	} else {
		fmt.Println(1)
	}

}

func ConnRedis(conn_str string, db int64) error {
	RPool = &redis.Pool{
		MaxIdle:     1,
		IdleTimeout: 60 * time.Second,
		Dial: func() (redis.Conn, error) {
			//c, err := redis.Dial("tcp", conn_str, redis.DialConnectTimeout(time.Duration(10)*time.Second),
			//	redis.DialReadTimeout(time.Duration(100)*time.Second),
			//	redis.DialWriteTimeout(time.Duration(100)*time.Second))
			c, err := redis.Dial("tcp", conn_str)
			if err != nil {
				return nil, err
			}
			if _, err := c.Do("SELECT", db); err != nil {
				c.Close()
				return nil, err
			}
			return c, nil
		},
	}
	return RPool.Get().Err()
}
