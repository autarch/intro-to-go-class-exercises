// I wrote this so I had a dead simple implementation of the tool to double
// check my conccurent implementations against.
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	words := map[string]int{}

	for _, n := range []string{"words1.txt", "words2.txt", "words3.txt"} {
		countWords(n, words)
	}

	keys := []string{}
	for k, _ := range words {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		fmt.Printf("%s - %d\n", k, words[k])
	}
}

func countWords(file string, words map[string]int) {
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}

	s := bufio.NewScanner(f)
	for s.Scan() {
		words[s.Text()]++
	}
}
