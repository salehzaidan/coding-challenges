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
			name: "single file",
			args: []string{"../test.txt"},
			in:   "",
			out: `"Your heart is the size of an ocean. Go find yourself in its hidden depths."
"The Bay of Bengal is hit frequently by cyclones. The months of November and May, in particular, are dangerous in this regard."
"Thinking is the capital, Enterprise is the way, Hard Work is the solution."
"If You Can'T Make It Good, At Least Make It Look Good."
"Heart be brave. If you cannot be brave, just go. Love's glory is not a small thing."
"It is bad for a young man to sin; but it is worse for an old man to sin."
"If You Are Out To Describe The Truth, Leave Elegance To The Tailor."
"O man you are busy working for the world, and the world is busy trying to turn you out."
"While children are struggling to be unique, the world around them is trying all means to make them look like everybody else."
"These Capitalists Generally Act Harmoniously And In Concert, To Fleece The People."`,
			err: nil,
		},
		{
			name: "stdin",
			args: nil,
			in:   `"Life isn’t about getting and having, it’s about giving and being."`,
			out:  `"Life isn’t about getting and having, it’s about giving and being."`,
			err:  nil,
		},
		{
			name: "multiple files",
			args: []string{"../test.txt", "../test2.txt"},
			in:   "",
			out: `"Your heart is the size of an ocean. Go find yourself in its hidden depths."
"The Bay of Bengal is hit frequently by cyclones. The months of November and May, in particular, are dangerous in this regard."
"Thinking is the capital, Enterprise is the way, Hard Work is the solution."
"If You Can'T Make It Good, At Least Make It Look Good."
"Heart be brave. If you cannot be brave, just go. Love's glory is not a small thing."
"It is bad for a young man to sin; but it is worse for an old man to sin."
"If You Are Out To Describe The Truth, Leave Elegance To The Tailor."
"O man you are busy working for the world, and the world is busy trying to turn you out."
"While children are struggling to be unique, the world around them is trying all means to make them look like everybody else."
"These Capitalists Generally Act Harmoniously And In Concert, To Fleece The People."
"I Don'T Believe In Failure. It Is Not Failure If You Enjoyed The Process."
"Do not get elated at any victory, for all such victory is subject to the will of God."
"Wear gratitude like a cloak and it will feed every corner of your life."
"If you even dream of beating me you'd better wake up and apologize."
"I Will Praise Any Man That Will Praise Me."
"One Of The Greatest Diseases Is To Be Nobody To Anybody."
"I'm so fast that last night I turned off the light switch in my hotel room and was in bed before the room was dark."
"People Must Learn To Hate And If They Can Learn To Hate, They Can Be Taught To Love."
"Everyone has been made for some particular work, and the desire for that work has been put in every heart."
"The less of the World, the freer you live."`,
			err: nil,
		},
		{
			name: "with line numbers",
			args: []string{"-n"},
			in: `"Life isn’t about getting and having, it’s about giving and being."
"Whatever the mind of man can conceive and believe, it can achieve."
"Strive not to be a success, but rather to be of value."`,
			out: `     1  "Life isn’t about getting and having, it’s about giving and being."
     2  "Whatever the mind of man can conceive and believe, it can achieve."
     3  "Strive not to be a success, but rather to be of value."`,
			err: nil,
		},
		{
			name: "with line numbers on blank/non-blank lines",
			args: []string{"-n"},
			in: `"Life isn’t about getting and having, it’s about giving and being."

"Whatever the mind of man can conceive and believe, it can achieve."

`,
			out: `     1  "Life isn’t about getting and having, it’s about giving and being."
     2  
     3  "Whatever the mind of man can conceive and believe, it can achieve."
     4  `,
			err: nil,
		},
		{
			name: "with line numbers on non-blank lines",
			args: []string{"-b"},
			in: `"Life isn’t about getting and having, it’s about giving and being."

"Whatever the mind of man can conceive and believe, it can achieve."

"Strive not to be a success, but rather to be of value."
`,
			out: `     1  "Life isn’t about getting and having, it’s about giving and being."

     2  "Whatever the mind of man can conceive and believe, it can achieve."

     3  "Strive not to be a success, but rather to be of value."`,
			err: nil,
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
