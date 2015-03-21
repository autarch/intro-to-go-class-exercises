package testprog

import (
	"io/ioutil"
	"os/exec"
	"path"
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

		ran := path.Base(program) + " " + strings.Join(c.Args, " ")

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
			assert.Fail("Ran '" + ran + "' and got an unexpected error: " + string(error))
		} else {
			if c.Stderr != nil {
				if err != nil {
					assert.Regexp(c.Stderr, string(error),
						"stderr from "+program+" matches \""+c.Stderr.String()+"\"")
				} else {
					assert.Fail("Ran '" + ran + "' and expected an error matching \"" +
						c.Stderr.String() +
						"\" but there was no error (stdout was \"" + string(output) + "\")")
				}
			} else if len(c.Stdout) > 0 {
				o := string(output)
				if trim {
					o = strings.TrimSpace(o)
				}
				assert.Equal(c.Stdout, o,
					"stdout from '"+ran+"' should be \""+c.Stdout+"\"")
			} else {
				t.Fatal("test case without stdout or stderr!")
			}
		}
	}
}
