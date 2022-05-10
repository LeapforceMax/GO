package main

import (
	"fmt"
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {

	m := make(map[string]int)

	f := strings.Fields(s)

	for _, value := range f {
		m[value]++
	}

	fmt.Println(m)

	return map[string]int{"x": 1}
}

func main() {
	wc.Test(WordCount)
}
