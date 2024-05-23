package cmd

import (
	"io"
	"os"

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

	filename := args[0]
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	if flagBytes {
		byteCounts, err := file.Seek(0, io.SeekEnd)
		if err != nil {
			return err
		}

		cmd.Printf("%d %s\n", byteCounts, filename)
	}

	return nil
}

func rootCmdFlags(cmd *cobra.Command) {
	cmd.Flags().BoolP("bytes", "c", false, "print the byte counts")
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	rootCmdFlags(rootCmd)
	rootCmd.SetOut(os.Stdout)
}
