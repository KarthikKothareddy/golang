package main

import (
	"fmt"
)

func main() {
	// arrays()
	// slices()
	sliceOperations()
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

	// looping over a slice
	fmt.Println(len(x))

	// loop by using range
	for i, v := range x {
		fmt.Println(i, v)
	}
	// loop without range
	for i := 0; i <= 4; i++ {
		fmt.Println(i, x[i])
	}

	// slicing a slice by colon
	// [start : end] where start is INCLUSIVE and end is EXCLUSIVE
	fmt.Println(x[1:3])

}

func sliceOperations() {

	x := []int{5, 6, 7, 8, 9}
	fmt.Println(x)

	// append to a slice
	x = append(x, 11, 22, 33)
	fmt.Println(x)

	y := []int{555, 666, 777}
	x = append(x, y...)

	fmt.Println(x)

}
