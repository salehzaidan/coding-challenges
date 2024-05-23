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
	}

	root := &cobra.Command{RunE: rootCmdRunE}
	rootCmdFlags(root)

	for _, tc := range tt {
		out, err := execute(t, root, tc.in, tc.args...)

		t.Run(tc.name, func(t *testing.T) {
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
