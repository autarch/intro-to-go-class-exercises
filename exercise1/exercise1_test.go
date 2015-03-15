package main

import (
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"regexp"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExercise1(t *testing.T) {
	dir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	var binary string
	if path.Base(dir) == "exercise1" {
		binary = "exercise1"
	} else {
		binary = "answer"
	}

	absBinary := path.Join(dir, binary)

	_, err = os.Open(absBinary)
	if err != nil {
		t.Fatal("Looks like there is no " + binary + " binary here. Did you run go build?")
	}

	t.Log("Found a binary named " + binary)

	testBinary(t, absBinary)
}

type testcase struct {
	args   []string
	stdout string
	stderr *regexp.Regexp
}

var cases = []testcase{
	testcase{
		args:   []string{"2", "4", "+"},
		stdout: "6",
		stderr: nil,
	},
	testcase{
		args:   []string{"42", "9", "+"},
		stdout: "51",
		stderr: nil,
	},
	testcase{
		args:   []string{"123456789", "123456790", "+"},
		stdout: "246913579",
		stderr: nil,
	},
	testcase{
		args:   []string{"-42", "1", "+"},
		stdout: "-41",
		stderr: nil,
	},
	testcase{
		args:   []string{"-42", "-5", "+"},
		stdout: "-47",
		stderr: nil,
	},
	testcase{
		args:   []string{"-42", "+"},
		stdout: "",
		stderr: regexp.MustCompile("This program expects 3 arguments - two numbers and an operator"),
	},
	testcase{
		args:   []string{},
		stdout: "",
		stderr: regexp.MustCompile("This program expects 3 arguments - two numbers and an operator"),
	},
	testcase{
		args:   []string{"1", "2", "3", "+"},
		stdout: "",
		stderr: regexp.MustCompile("This program expects 3 arguments - two numbers and an operator"),
	},
	testcase{
		args:   []string{"1", "2", "-"},
		stdout: "",
		stderr: regexp.MustCompile("Unknown operator: -"),
	},
}

func testBinary(t *testing.T, binary string) {
	assert := assert.New(t)

	for _, c := range cases {
		cmd := exec.Command(binary, c.args...)

		ran := path.Base(binary) + " " + strings.Join(c.args, " ")

		stdout, err := cmd.StdoutPipe()
		if err != nil {
			t.Fatalf("%s: %s", ran, err)
		}

		stderr, err := cmd.StderrPipe()
		if err != nil {
			t.Fatalf("%s: %s", ran, err)
		}

		if err = cmd.Start(); err != nil {
			t.Fatalf("%s: %s", ran, err)
		}

		output, err := ioutil.ReadAll(stdout)
		if err != nil {
			t.Fatalf("%s: %s", ran, err)
		}

		error, err := ioutil.ReadAll(stderr)
		if err != nil {
			t.Fatalf("%s: %s", ran, err)
		}

		err = cmd.Wait()

		if err != nil && c.stderr == nil {
			assert.Fail("Ran '" + ran + "' and got an unexpected error: " + string(error))
		} else {
			if c.stderr != nil {
				if err != nil {
					assert.Regexp(c.stderr, string(error),
						"stderr from "+binary+" matches \""+c.stderr.String()+"\"")
				} else {
					assert.Fail("Ran '" + ran + "' and expected an error matching \"" +
						c.stderr.String() +
						"\" but there was no error (stdout was \"" + string(output) + "\")")
				}
			} else if len(c.stdout) > 0 {
				o := strings.TrimSpace(string(output))
				assert.Equal(c.stdout, o,
					"stdout from '"+ran+"' should be \""+c.stdout+"\"")
			} else {
				t.Fatal("test case without stdout or stderr!")
			}
		}
	}
}
