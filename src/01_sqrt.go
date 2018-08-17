package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	z := 1.0
	for ; z < 100.0; z+=1.0 {
		res := (z*z - x) / (2*z)
		if res > 1 { break }
		fmt.Println(z, res)
	}
	return z
}

func main() {
	fmt.Println(Sqrt(17))
}

