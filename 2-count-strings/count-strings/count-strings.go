/*

Write a program that takes a list of strings on the command line and counts
the occurrence of each string. Then output the strings and their counts like
this:

bar - 1
foo - 2
quux - 1

You can use fmt.Printf() to print these out with a format of "%s - %d\n"

The output should be sorted by the strings (not their counts). You'll need to
use the core sort package for sorting.

If no strings are given at all, log a fatal error with this message:

    This program expects at least 1 argument

Note that Go does not let you declare a map variable and then start assigning
to it:

    var foo map[int]float
    foo[42] = 42.0

This will blow up with an error like "assignment to entry in nil map"

You can either use make() or create an empty map:

    foo := make(map[int]float)
    // The {} creates an empty (but initialized) map
    foo := map[int]float{}

To run the tests, run "go build", then "go test -v".

    > go build && go test -v

You're done when all the tests pass.

*/
package main

func main() {

}
