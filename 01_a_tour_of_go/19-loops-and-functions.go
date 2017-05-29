package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	for i := 0; i < 10; i++ {
		diff := ( math.Pow(z, 2) - x ) / ( 2 * z )
		if math.Abs(diff) < 0.00000001 {
			fmt.Println(i)
			return z
		} else {
			z = z - diff
		}
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(math.Sqrt(2))
}
