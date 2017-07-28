package db

import (
	"time"

	"github.com/garyburd/redigo/redis"
)

var RPool *redis.Pool

func ConnRedis(conn_str string) error {

	var err error
	RPool = &redis.Pool{
		MaxIdle:     1,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			var c redis.Conn
			c, err = redis.Dial("tcp", "ap2.jd.local:5360")
			if err != nil {
				return nil, err
			}
			if _, err := c.Do("AUTH", conn_str); err != nil {
				c.Close()
				return nil, err
			}
			return c, nil
		},
	}
	return err
}
