package main

import (
	"fmt"
	"math"
)

const threshold float64 = 1e-10

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %f\n", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0.0 {
		return x, ErrNegativeSqrt(x)
	}

	z := 1.0
	prev := z
	for i := 0; ; i++ {
		z -= (z*z - x) / (2 * z)
		//fmt.Println(i, z, math.Abs(prev - z))
		if math.Abs(prev-z) < threshold {
			break
		}
		prev = z
	}
	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
