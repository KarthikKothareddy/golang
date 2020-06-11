package main

import (
	"fmt"
)

func main() {
	// arrays()
	slices()
}

// GoLang recommends usage of SLICES than ARRAYS
func arrays() {

	// array is a building block in go
	// arrays have ZERO based indexes
	var x [5]int
	fmt.Println(x)

	x[2] = 11
	fmt.Println(x)

	// get the length of array
	fmt.Println(len(x))

}

func slices() {

	// composite literal
	// x := type{values}
	// A Slice allows you to group the values of the SAME type
	x := []int{1, 2, 3, 4, 5}
	fmt.Println(x)

}
