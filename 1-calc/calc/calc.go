/*

Write a very simple calculator program that takes three arguments, two numbers
and then the operation. The numbers will always be integers.

   > ./calc 2 4 +
   6
   > ./calc 4 5 "*"
   20

Note that the the shell will expand an asterisk (*) in place, so you need to
quote it when passing it on the command line.

The arguments will be available in the os.Args variable, which is a string
slice ([]string).

Parse the numbers into int64 values using strconv.ParseInt() function. See
http://golang.org/pkg/strconv/#ParseInt for documentation. If parsing
fails, log an error using log.Fatal(). The content of the error can be
anything you want (this is not tested).

If the program is given the wrong number of arguments, log a fatal error using
log.Fatal() with this message:

   This program expects 3 arguments - two numbers and an operator

The program should only handle two operators, addition (+) and multiplication
(*). If the program receives an operator it doesn't expect, log a fatal error
with this message:

   Unknown operator: {operator}

Where "{operator}" is replaced by the actual operator the program received.

To run the tests, run "go build", then "go test -v".

    > go build && go test -v

You can also run your program directly to manually test it:

    > go build && ./calc 2 3 +

You're done when all the tests pass.

*/
package main

func main() {

}
