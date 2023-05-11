package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	var z = 1.0
	var old float64
	for math.Abs(old-z) >= 0.000001 {
		old = z
		z -= (z*z - x) / (2 * z)
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2), math.Sqrt(2))
	fmt.Println(Sqrt(10), math.Sqrt(10))
}
