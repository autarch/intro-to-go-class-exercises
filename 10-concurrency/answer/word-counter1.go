package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

func main() {
	words := make(map[string]int)
	done := make(chan bool)
	c1 := makeReader("words1.txt", done)
	c2 := makeReader("words2.txt", done)
	c3 := makeReader("words3.txt", done)

	i := 0
	for i < 3 {
		select {
		case w := <-c1:
			words[w]++
		case w := <-c2:
			words[w]++
		case w := <-c3:
			words[w]++
		case <-done:
			i++
		}
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

func makeReader(file string, done chan<- bool) <-chan string {
	c := make(chan string)
	go func() {
		f, err := os.Open(file)
		if err != nil {
			log.Fatal(err)
		}

		s := bufio.NewScanner(f)
		for s.Scan() {
			c <- s.Text()
		}
		done <- true
	}()
	return c
}
