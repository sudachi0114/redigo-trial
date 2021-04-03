package main

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

func main() {
	fmt.Println("Redigo trial.")
	c := Connection()
	defer c.Close()
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
