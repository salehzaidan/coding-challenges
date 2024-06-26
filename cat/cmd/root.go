package cmd

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:  "cccat [file]...",
	RunE: rootCmdRunE,
}

func rootCmdRunE(cmd *cobra.Command, args []string) error {
	flagNumber, err := cmd.Flags().GetBool("number")
	if err != nil {
		return err
	}
	flagNumberNonBlank, err := cmd.Flags().GetBool("number-nonblank")
	if err != nil {
		return err
	}

	if flagNumberNonBlank {
		flagNumber = false
	}

	var filenames []string
	if len(args) >= 1 {
		filenames = args
	} else {
		filenames = []string{""}
	}

	i := 0
	lines := make([]string, 0)
	for _, filename := range filenames {
		var input io.Reader
		if filename == "" || filename == "-" {
			input = cmd.InOrStdin()
		} else {
			file, err := os.Open(filename)
			if err != nil {
				return err
			}
			defer file.Close()
			input = file
		}

		scanner := bufio.NewScanner(input)
		for scanner.Scan() {
			if flagNumberNonBlank && scanner.Text() != "" || flagNumber {
				lines = append(lines, fmt.Sprintf("%6d  %s\n", i+1, scanner.Text()))
				i++
			} else {
				lines = append(lines, fmt.Sprintf("%s\n", scanner.Text()))
			}
		}
	}

	cmd.Printf("%s", strings.Join(lines, ""))
	return nil
}

func rootCmdFlags(cmd *cobra.Command) {
	cmd.Flags().BoolP("number", "n", false, "number all output lines")
	cmd.Flags().BoolP("number-nonblank", "b", false, "number nonempty output lines, overrides -n")
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	rootCmdFlags(rootCmd)
	rootCmd.SetOut(os.Stdout)
}
