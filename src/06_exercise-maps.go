package main

import (
	"strings"
	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	
	var ret = make(map[string]int)
	
	for _, k := range strings.Fields(s) {
		ret[k] += 1
	}
	
	return ret
}

func main() {
	wc.Test(WordCount)
}
