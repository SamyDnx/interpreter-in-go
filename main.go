package main

import (
	"fmt"
	"monkey/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! Welcome the the Monkey programming language!\n", user.Username)
	fmt.Printf("Feel Free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
