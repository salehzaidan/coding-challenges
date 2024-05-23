package cmd

import (
	"bytes"
	"strings"
	"testing"

	"github.com/spf13/cobra"
)

func execute(t *testing.T, c *cobra.Command, in string, args ...string) (string, error) {
	t.Helper()

	if in != "" {
		c.SetIn(strings.NewReader(in))
	}

	buf := new(bytes.Buffer)
	c.SetOut(buf)
	c.SetErr(buf)
	c.SetArgs(args)

	err := c.Execute()
	return strings.TrimRight(buf.String(), "\n"), err
}

func TestRootCmd(t *testing.T) {
	tt := []struct {
		name string
		args []string
		in   string
		out  string
		err  error
	}{
		{
			name: "print byte counts",
			args: []string{"-c", "../test.txt"},
			in:   "",
			out:  "342190 ../test.txt",
			err:  nil,
		},
		{
			name: "print newline counts",
			args: []string{"-l", "../test.txt"},
			in:   "",
			out:  "7145 ../test.txt",
			err:  nil,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			root := &cobra.Command{RunE: rootCmdRunE}
			rootCmdFlags(root)
			out, err := execute(t, root, tc.in, tc.args...)

			if err != tc.err {
				t.Errorf("expect %q, got %q", tc.err, err)
				return
			}

			if tc.err == nil && out != tc.out {
				t.Errorf("expect %q, got %q", tc.out, out)
				return
			}
		})
	}
}
