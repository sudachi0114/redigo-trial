package infra

import (
	"fmt"
	"strings"

	"github.com/gomodule/redigo/redis"
)

func Login(message string, conn redis.Conn) (string, string) {
	splited := strings.Split(message, " ")
	username := splited[1]

	// Login したら GET で登録されているかを確認 => TTL を有効にして上書き
	//	しようと思ったのだけど、GETSET で TTL は登録できないらしい。
	//		login が登録とログインを兼ねてしまっていてちょっと気持ち悪い..
	//		だったら create 1つでいいのでは..??
	userkey := "online." + username
	res, err := conn.Do("SET", userkey, username, "NX", "EX", "120")
	if err != nil {
		panic(err)
	}

	if res == nil {
		fmt.Printf("User %s has already online.\n", username)
	} else {
		fmt.Printf("Logined User:[%s] | status:[%s].\n", username, res)
		fmt.Println(username)
	}

	return username, userkey
}
