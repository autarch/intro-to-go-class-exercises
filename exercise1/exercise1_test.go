package main

import (
	"regexp"
	"testing"

	"github.com/autarch/intro-to-go-class-exercises/helpers/testprog"
)

var cases = []testprog.TestCase{
	testprog.TestCase{
		Args:   []string{"2", "4", "+"},
		Stdout: "6",
		Stderr: nil,
	},
	testprog.TestCase{
		Args:   []string{"42", "9", "+"},
		Stdout: "51",
		Stderr: nil,
	},
	testprog.TestCase{
		Args:   []string{"123456789", "123456790", "+"},
		Stdout: "246913579",
		Stderr: nil,
	},
	testprog.TestCase{
		Args:   []string{"-42", "1", "+"},
		Stdout: "-41",
		Stderr: nil,
	},
	testprog.TestCase{
		Args:   []string{"-42", "-5", "+"},
		Stdout: "-47",
		Stderr: nil,
	},
	testprog.TestCase{
		Args:   []string{"-42", "+"},
		Stdout: "",
		Stderr: regexp.MustCompile("This program expects 3 arguments - two numbers and an operator"),
	},
	testprog.TestCase{
		Args:   []string{},
		Stdout: "",
		Stderr: regexp.MustCompile("This program expects 3 arguments - two numbers and an operator"),
	},
	testprog.TestCase{
		Args:   []string{"1", "2", "3", "+"},
		Stdout: "",
		Stderr: regexp.MustCompile("This program expects 3 arguments - two numbers and an operator"),
	},
	testprog.TestCase{
		Args:   []string{"1", "2", "-"},
		Stdout: "",
		Stderr: regexp.MustCompile("Unknown operator: -"),
	},
}

func TestExercise1(t *testing.T) {
	testprog.TestProgram(t, cases, true)
}
