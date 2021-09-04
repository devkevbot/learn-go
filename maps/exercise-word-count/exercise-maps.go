package main

import (
	"fmt"
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {

	var m map[string]int
	m = make(map[string]int)

	words := strings.Fields(s)
	fmt.Println(words)

	for _, word := range words {
		m[word]++
	}
	return m
}

func main() {
	wc.Test(WordCount)
}
