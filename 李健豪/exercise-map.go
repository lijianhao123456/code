package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	s1 := strings.Fields(s)
	m := make(map[string]int)
	for i := 0; i < len(s1); i++ {
		m[s1[i]]++
	}
	return m
}

func main() {
	wc.Test(WordCount)
}
