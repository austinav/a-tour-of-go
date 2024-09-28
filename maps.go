package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	result := make(map[string]int)
	for _, w := range strings.Fields(s) {
		_, ok := result[w]
		if ok {
			result[w]++
		} else {
			result[w] = 1
		}
	}

	return result
}

func main() {
	wc.Test(WordCount)
}
