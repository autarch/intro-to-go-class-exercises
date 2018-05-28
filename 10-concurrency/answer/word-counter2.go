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
	c1 := makeReader("words1.txt")
	c2 := makeReader("words2.txt")
	c3 := makeReader("words3.txt")

	i := 0
	for i < 3 {
		select {
		case w := <-c1:
			merge(w, words)
			i++
		case w := <-c2:
			merge(w, words)
			i++
		case w := <-c3:
			merge(w, words)
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

func makeReader(file string) <-chan map[string]int {
	c := make(chan map[string]int)
	go func() {
		f, err := os.Open(file)
		if err != nil {
			log.Fatal(err)
		}

		words := map[string]int{}
		s := bufio.NewScanner(f)
		for s.Scan() {
			words[s.Text()]++
		}
		c <- words
	}()
	return c
}

func merge(w1, w2 map[string]int) {
	for k, v := range w1 {
		w2[k] += v
	}
}
