package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/sudachi0114/redigo-trial/infra"
)

var (
	username string
	userkey  string
)

func main() {
	conn := infra.Connection()
	defer conn.Close()

	tickerChan := time.NewTicker(time.Millisecond * 1000).C

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

	chatExit := false

	for !chatExit {
		select {
		case <-tickerChan:
		case maybeMessage := <-sayChan:
			if maybeMessage == ".exit" {
				chatExit = true
			} else if strings.Contains(maybeMessage, ".create") {
				infra.CreateUser(maybeMessage, conn)
			} else if strings.Contains(maybeMessage, ".login") {
				username, userkey = infra.Login(maybeMessage, conn)
				fmt.Println("return:", username, userkey)
			} else if maybeMessage == ".whoami" {
				fmt.Printf("%s (key: %s)", username, userkey)
			} else {
				fmt.Println(maybeMessage)
			}
		}
	}
}
