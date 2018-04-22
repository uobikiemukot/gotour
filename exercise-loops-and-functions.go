package main

import (
	"fmt"
	"math"
)

const threshold = 1.0e-10

func Sqrt1(x float64) float64 {
	z := 5314532145321452314.0
	prev := z
	for i := 0; ; i++ {
		z -= (z*z - x) / (2 * z)
		fmt.Println(i, z, math.Abs(prev-z))
		if math.Abs(prev-z) < threshold {
			break
		}
		prev = z
	}
	return z
}

func main() {
	fmt.Println(Sqrt1(2))
	fmt.Println(math.Sqrt(2))
}
