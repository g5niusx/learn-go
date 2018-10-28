package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func main() {
	pool := &redis.Pool{Dial: func() (redis.Conn, error) {
		return redis.Dial("tcp", "127.0.0.1:6379")
	}}
	conn := pool.Get()
	defer conn.Close()
	conn.Do("set", "key", 9999)
	i, _ := redis.Int(conn.Do("get", "key"))
	fmt.Println("获取到的值:", i)

}
