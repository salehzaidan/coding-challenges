package cmd

import (
	"bufio"
	"io"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:  "ccwc [file]",
	Args: cobra.ExactArgs(1),
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

	filename := args[0]
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	var (
		byteCounts, lineCounts, wordCounts int64
	)

	if flagBytes {
		byteCounts, err = file.Seek(0, io.SeekEnd)
		if err != nil {
			return err
		}
		_, err := file.Seek(0, io.SeekStart)
		if err != nil {
			return err
		}
		cmd.Printf("%d ", byteCounts)
	}

	if flagLines || flagWords {
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			if flagLines {
				lineCounts++
			}

			if flagWords {
				wordCounts += int64(len(strings.Fields(scanner.Text())))
			}
		}

		if flagLines {
			cmd.Printf("%d ", lineCounts)
		}
		if flagWords {
			cmd.Printf("%d ", wordCounts)
		}
	}

	cmd.Printf("%s\n", filename)
	return nil
}

func rootCmdFlags(cmd *cobra.Command) {
	cmd.Flags().BoolP("bytes", "c", false, "print the byte counts")
	cmd.Flags().BoolP("lines", "l", false, "print the newline counts")
	cmd.Flags().BoolP("words", "w", false, "print the word counts")
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	rootCmdFlags(rootCmd)
	rootCmd.SetOut(os.Stdout)
}
