package main

import "fmt"

var s = []int{0, 1}

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func(int) []int {

	return func(x int) []int {
		var s_sli = s[(len(s)-2):];
		s = append(s, s_sli[0] + s_sli[1])
		return s
	}
	
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f(i))
	}
}

// go run awesomeProject/src/07_Successione_di_Fibonacci.go