package cmd

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:  "ccwc [file]",
	RunE: rootCmdRunE,
}

func rootCmdRunE(cmd *cobra.Command, args []string) error {
	flagBytes, err := cmd.Flags().GetBool("bytes")
	if err != nil {
		return err
	}
	flagLines, err := cmd.Flags().GetBool("lines")
	if err != nil {
		return err
	}
	flagWords, err := cmd.Flags().GetBool("words")
	if err != nil {
		return err
	}
	flagChars, err := cmd.Flags().GetBool("chars")
	if err != nil {
		return err
	}

	if !flagBytes && !flagLines && !flagWords && !flagChars {
		flagBytes = true
		flagLines = true
		flagWords = true
	}

	var filename string
	if len(args) >= 1 {
		filename = args[0]
	} else {
		filename = ""
	}

	var input io.ReadSeeker
	if filename == "" || filename == "-" {
		content, err := io.ReadAll(cmd.InOrStdin())
		if err != nil {
			return err
		}
		input = bytes.NewReader(content)
	} else {
		file, err := os.Open(filename)
		if err != nil {
			return err
		}
		defer file.Close()
		input = file
	}

	var (
		byteCounts, lineCounts, wordCounts, charCounts int
		maxCountLen                                    int
	)

	if flagLines || flagWords {
		scanner := bufio.NewScanner(input)
		for scanner.Scan() {
			if flagLines {
				lineCounts++
			}

			if flagWords {
				wordCounts += len(strings.Fields(scanner.Text()))
			}
		}
		_, err = input.Seek(0, io.SeekStart)
		if err != nil {
			return err
		}

		if flagLines {
			maxCountLen = max(maxCountLen, len(strconv.Itoa(lineCounts)))
		}
		if flagWords {
			maxCountLen = max(maxCountLen, len(strconv.Itoa(wordCounts)))
		}
	}

	if flagChars {
		content, err := io.ReadAll(input)
		if err != nil {
			return err
		}
		_, err = input.Seek(0, io.SeekStart)
		if err != nil {
			return err
		}
		charCounts = utf8.RuneCount(content)
		maxCountLen = max(maxCountLen, len(strconv.Itoa(charCounts)))
	}

	if flagBytes {
		n, err := input.Seek(0, io.SeekEnd)
		if err != nil {
			return err
		}
		byteCounts = int(n)
		_, err = input.Seek(0, io.SeekStart)
		if err != nil {
			return err
		}
		maxCountLen = max(maxCountLen, len(strconv.Itoa(byteCounts)))
	}

	countFormat := fmt.Sprintf("%%%dd ", maxCountLen)
	if flagLines {
		cmd.Printf(countFormat, lineCounts)
	}
	if flagWords {
		cmd.Printf(countFormat, wordCounts)
	}
	if flagChars {
		cmd.Printf(countFormat, charCounts)
	}
	if flagBytes {
		cmd.Printf(countFormat, byteCounts)
	}
	cmd.Printf("%s\n", filename)
	return nil
}

func rootCmdFlags(cmd *cobra.Command) {
	cmd.Flags().BoolP("bytes", "c", false, "print the byte counts")
	cmd.Flags().BoolP("lines", "l", false, "print the newline counts")
	cmd.Flags().BoolP("words", "w", false, "print the word counts")
	cmd.Flags().BoolP("chars", "m", false, "print the character counts")
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	rootCmdFlags(rootCmd)
	rootCmd.SetOut(os.Stdout)
}
