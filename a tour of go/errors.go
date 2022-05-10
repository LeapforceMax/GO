package main

import (
	"errors"
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("Foutmelding!")
}

func Sqrt(x float64) (float64, error) {
	z := 2.0
	y := 0.0

	// Error handling kan ook op deze manier, Is een stuk gemakkelijker
	if x < 0 {
		return 0, errors.New("Cannot take sqrt from negative number")
	}

	for i := 0; i < 10; i++ {
		fmt.Println(math.Pow(z, 2.0))

		y = z - (z*z-x)/(2*z)
		//fmt.Println(y)

		if d := math.Abs(z - y); d < 0.05 {
			//fmt.Println(d)
			z = y
		} else {
			z = y
		}

	}
	return 0, nil
}

func main() {

	fmt.Println(Sqrt(-2))
}
