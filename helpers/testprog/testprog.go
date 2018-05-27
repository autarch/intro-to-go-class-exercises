package testprog

import (
	"io/ioutil"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
	"testing"

	"github.com/autarch/intro-to-go-class-exercises/helpers/findbin"
	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	Args   []string
	Stdout string
	Stderr *regexp.Regexp
}

func TestProgram(t *testing.T, cases []TestCase, trim bool) {
	program := findbin.FindBin(t)

	assert := assert.New(t)

	for _, c := range cases {
		cmd := exec.Command(program, c.Args...)

		ran := filepath.Base(program) + " " + strings.Join(c.Args, " ")

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

		if err != nil && c.Stderr == nil {
			assert.Fail(`Ran "%s" and got an unexpected error: %s`, error, ran)
		} else {
			if c.Stderr != nil {
				if err != nil {
					assert.Regexp(
						c.Stderr, string(error),
						`stderr from "%s" matches "%s"`,
						ran, c.Stderr.String(),
					)
				} else {
					assert.Fail(
						`Ran '%s' and expected an error matching "%s" but there was no error (stdout was "%s")`,
						ran, c.Stderr.String(), string(output),
					)
				}
			} else if len(c.Stdout) > 0 {
				o := string(output)
				if trim {
					o = strings.TrimSpace(o)
				}
				assert.Equal(
					c.Stdout, o,
					`stdout from "%s" should be "%s"`, ran, c.Stdout,
				)
			} else {
				t.Fatal("test case without stdout or stderr!")
			}
		}
	}
}
