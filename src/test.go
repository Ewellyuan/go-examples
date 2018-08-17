package main

import "fmt"

func main() {
	
	arr1 := []string{"Hello"}
	arr2 := []string{"Better", "Tomorrow"}

	slice1 := arr1[:]
	slice2 := append(slice1, arr2)

	fmt.Println(slice2)

}

func append(slice []string, data []string) []string {

	tmp_slice := make([]string, (len(slice) + len(data)), (cap(slice) + cap(data)))
	copy(tmp_slice, slice)

	for i:=0;i<len(data);i++ {
		tmp_slice[len(slice)+i] = data[i]
	}

	return tmp_slice

}