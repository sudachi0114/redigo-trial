package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/gomodule/redigo/redis"
	"github.com/sudachi0114/redigo-trial/infra"
)

var (
	username string
	userkey  string
)

func main() {
	conn := infra.Connection()
	defer conn.Close()

	tickerChan := time.NewTicker(time.Second * 60).C

	sayChan := make(chan string)
	go func() {
		prompt := "(屮`･д･)屮 "
		bio := bufio.NewReader(os.Stdin)
		for {
			fmt.Print(prompt)

			line, _, err := bio.ReadLine()
			if err != nil {
				fmt.Println(err)
				sayChan <- ".exit"
				return
			}
			sayChan <- string(line)
		}
	}()

	// PubSub の sub <- subscribe channel
	subChan := make(chan string)
	go func() {
		subConn := infra.Connection()
		defer subConn.Close()

		// pub/sub connection
		psc := redis.PubSubConn{Conn: subConn}
		psc.Subscribe("messages")
		for {
			switch v := psc.Receive().(type) {
			case redis.Message:
				subChan <- string(v.Data)
			case redis.Subscription:
				break
			case error:
				return
			}
		}
	}()

	chatExit := false

	for !chatExit {
		select {
		case <-tickerChan:
			if userkey != "" && username != "" {
				res, err := conn.Do("SET", userkey, username, "XX", "EX", "120")
				if err != nil || res == nil {
					fmt.Println("Heartbeat set failed..")
					chatExit = true
				}
			}
		case maybeMessage := <-sayChan:
			if maybeMessage == ".exit" {
				infra.ExitUser(userkey, username, conn)
				chatExit = true
			} else if strings.Contains(maybeMessage, ".create") {
				infra.CreateUser(maybeMessage, conn)
			} else if strings.Contains(maybeMessage, ".login") {
				username, userkey = infra.Login(maybeMessage, conn)
				fmt.Println("return:", username, userkey)
			} else if maybeMessage == ".whoami" {
				fmt.Printf("%s (key: %s)", username, userkey)
			} else {
				message := "[" + username + "] < " + maybeMessage
				conn.Do("PUBLISH", "messages", message)
			}
		case msg := <-subChan:
			fmt.Println(msg)
		}
	}
}
