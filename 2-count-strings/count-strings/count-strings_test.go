package main

import (
	"regexp"
	"testing"

	"github.com/autarch/intro-to-go-class-exercises/helpers/testprog"
)

var cases = []testprog.TestCase{
	{
		Args: []string{"foo", "bar", "quux", "zub", "bar", "foo", "a", "bar"},
		Stdout: `a - 1
bar - 3
foo - 2
quux - 1
zub - 1
`,
		Stderr: nil,
	},
	{
		Args: []string{"foo", "foo", "foo", "bar"},
		Stdout: `bar - 1
foo - 3
`,
		Stderr: nil,
	},
	{
		Args:   []string{},
		Stdout: "",
		Stderr: regexp.MustCompile("This program expects at least 1 argument"),
	},
}

func TestExercise2(t *testing.T) {
	testprog.TestProgram(t, cases, false)
}
