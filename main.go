package main

import (
	"boulang/repl"
	"fmt"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! Welcome to Boulang v0.0.1.\n", user.Username)
	fmt.Printf("Feel free to type in commands.\n")
	repl.Start(os.Stdin, os.Stdout)
}
