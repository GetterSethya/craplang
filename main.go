package main

import (
	"bufio"
	"crapLang/repl"
	"fmt"
	"os"
	"path/filepath"
)

func main() {

	args := os.Args
	if len(args) > 1 {
        // cek file extension apakah .crap
		fileExtensions := filepath.Ext(args[1])
		if fileExtensions != ".crap" {
			fmt.Printf("error: %+v is not supported\n", fileExtensions)
            return
		}
		// buka file
		// input ke std in
		file, err := os.Open(args[1])
		if err != nil {
			fmt.Printf("error: %+v\n", err)
			return
		}
		defer file.Close()
		reader := bufio.NewReader(file)
		repl.Start(reader, os.Stdout, false)
	} else {

		fmt.Printf("REPL 💩-lang\n")
		repl.Start(os.Stdin, os.Stdout, true)
	}

}
