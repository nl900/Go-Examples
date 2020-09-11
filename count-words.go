package main

import (
	"fmt"
	"golang.org/x/tour/wc"
	"strings"
)

/*
Counts occurrences of words in string
*/
func WordCount(s string) map[string]int {
	var counts map[string]int
	counts = make(map[string]int)
	for _, word := range strings.Fields(s) {
		_, ok := counts[word]
		if ok {
			counts[word] += 1
		} else {
			counts[word] = 1
		}
	
	}
	return counts
}

func main() {
	wc.Test(WordCount)
}
