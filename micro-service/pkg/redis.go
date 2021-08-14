package pkg

import (
	"github.com/gomodule/redigo/redis"
	"log"
)

func NewRedisClient () (c redis.Conn){
	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		log.Panicf("redis Dial() err:%v", err)
		return nil
	}
	return c
}
