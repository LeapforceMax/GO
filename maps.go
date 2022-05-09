package main

import (
	"fmt"
	"strings"
	"unicode"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {

	m := make(map[string]int)

	f := func(s rune) bool {
		return !unicode.IsLetter(s) && !unicode.IsNumber(s)
	}

	fmt.Println("Fields are: ", strings.FieldsFunc(s, f))

	return map[string]int{"x": 1}
}

func main() {
	wc.Test(WordCount)
}
