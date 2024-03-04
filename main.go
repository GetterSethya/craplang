package main

import (
	"crapLang/repl"
	"fmt"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
    if err != nil {
        panic(err)
    }

    fmt.Printf("Hello %s! this is the craplang programming language!\n", user.Username)
    fmt.Printf("Feel free to type in commands\n")
    repl.Start(os.Stdin, os.Stdout)
}
