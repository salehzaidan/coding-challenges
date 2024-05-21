package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	numberLinesEnabled := flag.Bool("n", false, "number all output lines")
	flag.Parse()

	var filenames []string
	if flag.NArg() >= 1 {
		filenames = flag.Args()
	} else {
		filenames = []string{""}
	}

	i := 0
	lines := make([]string, 0)
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

		scanner := bufio.NewScanner(input)
		for scanner.Scan() {
			if *numberLinesEnabled {
				lines = append(lines, fmt.Sprintf("%6d  %s\n", i+1, scanner.Text()))
			} else {
				lines = append(lines, fmt.Sprintf("%s\n", scanner.Text()))
			}
			i++
		}
	}

	fmt.Printf("%s", strings.Join(lines, ""))
}
