package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
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
		} else {
			fmt.Println(maybeMessage)
		}
	}
}
