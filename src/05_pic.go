package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {

	// init
	ret := make([][]uint8, dy)
	for i:=0; i<dy; i++ {
		ret[i] = make([]uint8, dx)
	}
	
	// draw
	for y,_ := range ret {
		for x,_ := range ret[y] {
			ret[y][x] = uint8(x ^ y)
		}
	}
	
	return ret
	
}

func main() {
	pic.Show(Pic)
}
