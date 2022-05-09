package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {

	z := 2.0
	y := 0.0

	for i := 0; i < 10; i++ {
		fmt.Println(math.Pow(z, 2.0))

		y = z - (z*z-x)/(2*z)
		//fmt.Println(y)

		if d := math.Abs(z - y); d < 0.05 {
			//fmt.Println(d)
			z = y
			return z
			break
		} else {
			z = y
		}

	}

	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
