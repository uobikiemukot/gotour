package main

import (
	"fmt"
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	//m := make(map[string]int)
	//return map[string]int{"x": 1}
	var m = make(map[string]int)

	fmt.Println(strings.Fields(s))

	substr := strings.Fields(s)
	for _, s := range substr {
		m[s]++
	}

	fmt.Println(m)

	return m
}

func main() {
	wc.Test(WordCount)
}
