package main

import (
	"crapLang/repl"
	"fmt"
	"os"
)

func main() {

	fmt.Printf("REPL 💩-lang\n")
	repl.Start(os.Stdin, os.Stdout)
}
