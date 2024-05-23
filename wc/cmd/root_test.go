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
		{
			name: "print word counts",
			args: []string{"-w", "../test.txt"},
			in:   "",
			out:  "58164 ../test.txt",
			err:  nil,
		},
		{
			name: "print character counts",
			args: []string{"-m", "../test.txt"},
			in:   "",
			out:  "339292 ../test.txt",
			err:  nil,
		},
		{
			name: "no flags",
			args: []string{"../test.txt"},
			in:   "",
			out:  "  7145  58164 342190 ../test.txt",
			err:  nil,
		},
		{
			name: "from stdin",
			args: nil,
			in: `The Project Gutenberg eBook of The Art of War

This ebook is for the use of anyone anywhere in the United States and
most other parts of the world at no cost and with almost no restrictions
whatsoever. You may copy it, give it away or re-use it under the terms
of the Project Gutenberg License included with this ebook or online
at www.gutenberg.org. If you are not located in the United States,
you will have to check the laws of the country where you are located
before using this eBook.
`,
			out: "  9  91 490 ",
			err: nil,
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
