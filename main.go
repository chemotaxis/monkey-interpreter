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

	fmt.Printf("Hi %s!  This is the Monkey language REPL.\n", user.Username)
	fmt.Printf("Feel free to type in commands.\n")
	fmt.Printf("To exit, type 'exit' or press ctrl+d or ctrl+c.\n")

	repl.Start(os.Stdin, os.Stdout)
}
