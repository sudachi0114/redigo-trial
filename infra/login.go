package infra

import (
	"fmt"
	"strings"

	"github.com/gomodule/redigo/redis"
)

func Login(message string, conn redis.Conn) (string, string) {
	splited := strings.Split(message, " ")
	username := splited[1]

	userkey, err := redis.String(conn.Do("GET", username)) // user が登録されているか?
	if err != nil {
		panic(err)
	}
	if userkey == "" {
		fmt.Printf("User %s is not exists.\n", username)
	}
	return username, userkey
}
