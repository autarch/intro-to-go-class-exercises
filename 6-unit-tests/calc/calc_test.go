/*

This exercise is a bit different. This time you're going to write tests for my
code, instead of writing code to be tested.

The package you're testing is the "calc" package. It implements the following
functions:

- Add(a, b int) int
- Sub(a, b int) int
- Mult(a, b int) int
- Div(a, b int) (float64, err)
- Sum(v ...int) int

The Div() function returns an error when asked to divide by 0. The text of the
error is "Cannot divide by 0." I suggest only testing with integers that
divide into a value with a finite precision (1/2, not 22/7), unless you enjoy
dealing with floating point precision issues.

The Sum function takes an arbitrary number of integers and returns their
sum. If given *no* arguments, it returns 0.

The Mult() function has a bug. Feel free to fix it after you've found it
through testing.

Feel free to use the assert package. See
http://godoc.org/github.com/stretchr/testify/assert for details.

Consider defining test cases in an array of arrays (or array of structs) and
iterating over them in order to avoid repetitive code.

I can't really test your tests, but ask me if you'd like me to take a look at
your test code and give you feedback.

*/
