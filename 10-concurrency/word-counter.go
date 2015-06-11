/*

The goal for this exercise is to open a set of files and count the words in
them. You should do this using goroutines to count the occurences of words in
each file concurrently.

The files are named "words1.txt", "words2.txt", and "words3.txt", and are
located in the current directory. Each file contains one word per line, with
no blank lines.

You should open each file and count the occurrences of words in the file. Your
ultimate output should be a list of words and the number of times they occur
across all the files. Sort the output alphabetically. All the words are lower
case, so you don't need to worry about this when sorting.

Use the os.Open() function to open the file, and then pass the resulting File
object to the bufio.NewScanner(). You can then call s.Scan() on the scanner in
a for loop with s.Text() to get the list of words in the file.

This will look something like this:

    f, err := os.Open(filename)
    if err != nil {
        panic(err)
    }
    s := bufio.NewScanner(f)
    for s.Scan() {
        word := s.Text()
    }

There are two approaches you can take to implementing concurrency here.

You could send *each word* to the main process via a channel and do all the
counting there using a map[string]int.

Alternatively, you could have your goroutines read the entire file and
increment the values of a map[string]int as you go. Then you can send the
entire map[string]int value to the main process via a channel. The main
process can then merge the counts.

The first implementation strategy is probably a bit easier, but if you have
time, try implementing both!

In order to sort the words, you'll need to place all the keys of the master
count map in an array, and then call sort.Strings() on that array. Then you
can iterate over that sorted array and print the word and its count.

The exact output you should see is in the expected.txt file. 

The easiest way to test your code is by running the following:

    go run ./word-counter.go > output.txt; diff -u expected.txt output.txt

If there are no differences between your output and the expected file then
your code is working.

*/
