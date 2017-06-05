package db

import "github.com/garyburd/redigo/redis"

var C redis.Conn

func ConnRedis() error {
	var err error
	C, err = redis.Dial("tcp", ":6379")
	if err != nil {
		return err
	}
	return nil
}

func StopRedis() {
	C.Close()
}
