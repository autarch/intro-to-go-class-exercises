/*

Write a very simply calculator program that takes three arguments, the
operation (+ or *), and then two numbers. The numbers will always be integers.

   > exercise1 2 4 +
   6
   > exercise1 4 5 *
   20

Scan the numbers into int64 values using the fmt package.

If the program is given the wrong number of arguments, log a fatal error using
log.Fatal() with this message:

   This program expects 3 arguments - two numbers and an operator

If the program receives an operator it doesn't expect, log a fatal error with
this message:

   Unknown operator: {operator}

Where "{operator}" is replaced by the actual operator the program received.

To run the tests, run "go build", then "go test -v".

*/
package main

func main() {

}
