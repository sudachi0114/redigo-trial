package infra

import "github.com/gomodule/redigo/redis"

func ExitUser(userkey, username string, conn redis.Conn) {
	conn.Do("DEL", userkey)
	conn.Do("SREM", "users", username)

	message := username + " has left."
	conn.Do("PUBLISH", "messages", message)
}
