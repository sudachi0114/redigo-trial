package infra

import (
	"github.com/gomodule/redigo/redis"
)

func Connection() redis.Conn {
	const Addr = "redis:6379"

	conn, err := redis.Dial("tcp", Addr)
	if err != nil {
		panic(err)
	}
	return conn
}
