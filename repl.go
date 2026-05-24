package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os/user"
	"strings"
)

const PROMPT = "%s@tinydb> "

func StartRepl(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	user, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	for {
		fmt.Printf(PROMPT, user.Username)
		if !scanner.Scan() {
			break
		}

		input := strings.TrimSpace(scanner.Text())

		if input == "\\quit" || input == "\\q" || input == "exit" {
			fmt.Println("bye")
			break
		}

	}
}
