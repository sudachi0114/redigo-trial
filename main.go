package main

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

func main() {
	fmt.Println("Redigo trial.")
	c := Connection()
	defer c.Close()

	// GET key
	res_get := Get("hoge", c)
	fmt.Println(res_get)
}

func Connection() redis.Conn {
	const Addr = "127.0.0.1:6379"

	c, err := redis.Dial("tcp", Addr)
	if err != nil {
		panic(err)
	}
	fmt.Println("Succesfly Connect to redis @", Addr)
	return c
}

func Get(key string, c redis.Conn) string {
	res, err := redis.String(c.Do("GET", key))
	if err != nil {
		panic(err)
	}
	return res
}
