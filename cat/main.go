package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	var filename string
	if len(os.Args) >= 2 {
		filename = os.Args[1]
	}

	var input io.Reader
	if filename == "" || filename == "-" {
		input = os.Stdin
	} else {
		file, err := os.Open(filename)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		defer file.Close()
		input = file
	}

	content, err := io.ReadAll(input)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Printf("%s", content)
}
