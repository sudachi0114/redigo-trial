package main

import (
	"fmt"
	"time"

	"github.com/sudachi0114/redigo-trial/infra"
)

func main() {
	conn := infra.Connection()
	defer conn.Close()

	// prompt := "(屮`･д･)屮 "
	// bio := bufio.NewReader(os.Stdin)
	// for {
	// 	fmt.Print(prompt)

	// 	line, _, err := bio.ReadLine()
	// 	maybeMessage := string(line)

	// 	if err != nil {
	// 		fmt.Println(err)
	// 		return
	// 	}

	// 	if maybeMessage == ".exit" {
	// 		return
	// 	} else if strings.Contains(maybeMessage, ".create") {
	// 		username, userkey := infra.CreateUser(maybeMessage, conn)
	// 		fmt.Println("return:", username, userkey)
	// 	} else {
	// 		fmt.Println(maybeMessage)
	// 	}
	// }

	tickerChan := time.NewTicker(time.Millisecond * 1000).C

	chatExit := false

	for !chatExit {
		select {
		case <-tickerChan:
			fmt.Println(" -- tickerChan -- ")
		}
	}

}
