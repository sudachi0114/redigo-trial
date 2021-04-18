package infra

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/gomodule/redigo/redis"
)

func CreateUser(message string, conn redis.Conn) {
	splited := strings.Split(message, " ")
	username := splited[1]

	userkey := username // FIXME: 果たしてこんなデータに意味はあるのだろうか..
	res, err := conn.Do("SET", userkey, username)
	if err != nil {
		panic(err) // redis.String(conn.Do()) にして、何回か同じユーザの登録をリクエストすると、ここでパニックする..??
	}

	if res == nil {
		fmt.Printf("username %s has already used.\n", username)
	} else {
		message := fmt.Sprintf("[system] Create User: %s | status: %s.", username, res)
		conn.Do("PUBLISH", "messages", message)
		// FIXME: AddtoUserGroup(username, conn)
	}
}

func AddtoUserGroup(user string, conn redis.Conn) {
	res, err := redis.String(conn.Do("SADD", "users", user))
	if err != nil {
		panic(err)
	}
	if reflect.TypeOf(res) == reflect.TypeOf(0) {
		fmt.Printf("user: %s has already in group: users\n", user)
	} else {
		fmt.Printf("Added user: %s to group: users\n", user)
	}
}
