package testbin

type Testcase struct {
	args   []string
	stdout string
	stderr *regexp.Regexp
}

func TestBinary(t *testing.T, binary string) {
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
