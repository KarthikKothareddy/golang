package main

import (
	"fmt"
)

func main() {
	// arrays()
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

}
