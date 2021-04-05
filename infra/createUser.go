package infra

import (
	"fmt"
	"strings"

	"github.com/gomodule/redigo/redis"
)

func CreateUser(message string, conn redis.Conn) {
	splited := strings.Split(message, " ")
	username := splited[1]

	userkey := "online." + username
	res, err := conn.Do("SET", userkey, username, "NX", "EX", "120")
	if err != nil {
		panic(err) // redis.String(conn.Do()) にして、何回か同じユーザの登録をリクエストすると、ここでパニックする..??
	}

	if res == nil {
		fmt.Printf("User %s has already online.\n", username)
	} else {
		fmt.Printf("Create User:[%s] | status:[%s].\n", username, res)
		AddtoUserGroup(username, conn)
	}
}

func AddtoUserGroup(user string, conn redis.Conn) string {
	res, err := redis.String(conn.Do("SADD", "users", user))
	if err != nil {
		panic(err)
	}
	fmt.Printf("Added user: %s to group: users\n", user)
	return res
}
