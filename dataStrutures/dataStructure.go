package main

import (
	"fmt"
)

func main() {
	// arrays()
	// slices()
	// sliceOperations()
	multiDimensionalSlice()
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

	// composite literal
	x := []int{5, 6, 7, 8, 9}
	fmt.Println(x)

	// append to a slice
	x = append(x, 11, 22, 33)
	fmt.Println(x)

	y := []int{555, 666, 777}
	x = append(x, y...)
	fmt.Println(x)

	// deleting from a slice
	// deleting values 7 and 8 from slice based on index
	x = append(x[:2], x[4:]...)
	fmt.Println(x)

	// SLICE MAKE
	// make(TYPE, INITIAL LENGTH, CAPACITY)
	z := make([]int, 10, 100)
	fmt.Println(z)
	fmt.Println(len(z))
	fmt.Println(cap(z))

	z = x
	fmt.Println(z)
	fmt.Println(len(z))
	fmt.Println(cap(z))

	z = append(z, 1024, 2048)
	fmt.Println(z)
	fmt.Println(len(z))
	fmt.Println(cap(z))

}

func multiDimensionalSlice() {

	p1 := []string{"James", "Bond", "Vanilla"}
	fmt.Println(p1)

	p2 := []string{"BananaSplit", "Chocolate", "CookiesNCream"}
	fmt.Println(p2)

	// multi dimensional slice
	p3 := [][]string{p1, p2}
	fmt.Println(p3)

}
