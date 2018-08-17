package main

import (
	"fmt"
	"math"
)

type Vertec struct {
	X, Y float64
}

func (v Vertec) Abs1() float64 {
	return math.Sqrt(v.X*v.X + v.Y+v.Y)
}

func Abs2(v Vertec) float64 {
	return math.Sqrt(v.X*v.X + v.Y+v.Y)
}

func main() {
	var v = Vertec{10, 20}
	fmt.Println(v.Abs1())
	fmt.Println(Abs2(v))
}