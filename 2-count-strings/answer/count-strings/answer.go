package main

import (
	"fmt"
	"log"
	"os"
	"sort"
)

func main() {
	if len(os.Args) == 1 {
		log.Fatal("This program expects at least 1 argument")
	}

	strings := map[string]int{}
	for _, s := range os.Args[1:] {
		if _, found := strings[s]; found {
			strings[s]++
		} else {
			strings[s] = 1
		}
	}

	keys := []string{}
	for k, _ := range strings {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	for _, s := range keys {
		fmt.Printf("%s - %d\n", s, strings[s])
	}
}
