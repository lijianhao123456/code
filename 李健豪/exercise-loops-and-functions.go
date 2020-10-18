package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	z1 := z
	for {
		z = z - (z*z-x)/(2*z)
		if math.Abs(z-z1) < 1e-15 {
			break
		}
		z1 = z
	}
	return z1
}

func main() {
	fmt.Println(Sqrt(3))
	fmt.Println(math.Sqrt(3))
}
