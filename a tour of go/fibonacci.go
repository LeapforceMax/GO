package main

import "fmt"

func fibonacci() func(int) int {
	return func(i int) int {
		var sequence int
		var n1, n2 int = 0, 1

		for x := 0; x <= i; x++ {
			sequence = n1 + n2
			n1 = n2
			n2 = sequence
		}

		return sequence
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f(i))
	}
}
