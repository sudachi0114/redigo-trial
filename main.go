package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/sudachi0114/redigo-trial/infra"
)

func main() {
	c := infra.Connection()
	defer c.Close()

	prompt := "(屮`･д･)屮 "
	bio := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(prompt)

		line, _, err := bio.ReadLine()
		maybeMessage := string(line)

		if err != nil {
			fmt.Println(err)
			return
		}

		if maybeMessage == ".exit" {
			return
		} else if strings.Contains(maybeMessage, ".create") {
			splited := strings.Split(maybeMessage, " ")
			userName := splited[1]
			fmt.Println("created user:", userName)
		} else {
			fmt.Println(maybeMessage)
		}
	}
}
