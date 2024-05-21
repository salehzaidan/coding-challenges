package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	var filenames []string
	if len(os.Args) >= 2 {
		filenames = os.Args[1:]
	} else {
		filenames = []string{""}
	}

	var output strings.Builder
	for _, filename := range filenames {
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
		output.WriteString(string(content))
	}

	fmt.Printf("%s", output.String())
}
