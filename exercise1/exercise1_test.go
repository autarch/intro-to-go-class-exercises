package main

import (
	"io/ioutil"
	"os"
	"os/exec"
	"path"
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

	_, err = os.Open(binary)
	if err != nil {
		t.Fatal("Looks like there is no " + binary + " binary here. Did you run go build?")
	}

	t.Log("Found a binary named " + binary)

	testBinary(t, binary)
}

type testcase struct {
	args   []string
	stdout string
	stderr *Regexp
}

func testBinary(t *testing.T, binary string) {
	cases := []testcase{
		testcase{
			args:   {"+", "2", "4"},
			stdout: "8",
			stderr: nil,
		},
		testcase{
			args:   {"+", "42", "9"},
			stdout: "51",
			stderr: nil,
		},
		testcase{
			args:   {"+", "123456789", "123456790"},
			stdout: "246913579",
			stderr: nil,
		},
		testcase{
			args:   {"+", "-42", "1"},
			stdout: "-41",
			stderr: nil,
		},
		testcase{
			args:   {"+", "-42", "-5"},
			stdout: "-47",
			stderr: nil,
		},
	}

	assert := assert.New(t)

	for _, c := range cases {
		cmd := os.Command(binary, c.args...)

		stdout, err := cmd.StdoutPipe()
		if err != nil {
			t.Fatal(err)
		}

		stderr, err := cmd.StdoutPipe()
		if err != nil {
			t.Fatal(err)
		}

		if err = cmd.Start(); err != nil {
			t.Fatal(err)
		}

		if output, err := ioutil.ReadAll(stdout); err != nil {
			t.Fatal(err)
		}

		if error, err := ioutil.ReadAll(stderr); err != nil {
			t.Fatal(err)
		}

		if err = cmd.Wait(); err != nil {
			t.Fatal(err)
		}

		if len(c.stdout) {
			assert.Equal(c.stdout, output, "stdout from "+binary+" is "+c.stdout)
		} else if c.stderr != nil {
			assert.Regexp(c.stderr, error, "stderr from "+binary+" matches "+c.stderr.String())
		}
	}
}
