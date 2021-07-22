package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := float64(2)
	n := float64(0)
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
		if math.Abs(z-n) < 1e-10 {
			break
		}
		n = z
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(math.Sqrt(2))
}
