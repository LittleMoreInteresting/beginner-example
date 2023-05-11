package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	res := make(map[string]int)
	words := strings.Fields(s)
	for _, word := range words {
		res[word]++
	}
	return res
}

func main() {
	wc.Test(WordCount)
}
