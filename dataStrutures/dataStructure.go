package main

import (
	"fmt"
)

func main() {
	// arrays()
	// slices()
	// sliceOperations()
	// multiDimensionalSlice()
	// maps()
	mapOperations()
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

func maps() {
	// composite literal
	// map[KeyType]valueType
	m := map[string]int{
		"James": 11,
		"Bond":  22,
	}
	fmt.Println(m)
	// accessing values by Key
	fmt.Println(m["James"])
	// defaults to 0
	fmt.Println(m["NonExistentKey"])

	// returns two values back
	v, ok := m["NonExistentKey"]
	fmt.Println(v, ok)

	// idiomatic GO Code
	if v, ok := m["Bond"]; ok {
		fmt.Println("The Bond Key exists: ", v)
	}

}

func mapOperations() {

	m := map[string]int{
		"James": 11,
		"Bond":  22,
	}

	// add an element to map
	m["Jason"] = 33

	// checking for a key in map
	if v, ok := m["Jason"]; ok {
		fmt.Println(v, ok)
	} else {
		fmt.Println("Jason Doesn't exist yet")
	}

	// looping through a map
	for k, v := range m {
		fmt.Println(k, v)
	}

	// range with slice
	xi := []int{1, 2, 3, 4, 5}

	for i, v := range xi {
		fmt.Println(i, v)
	}

	// deleting the key from map
	delete(m, "James")
	fmt.Println(m)

	// delete some key that doesn't exist
	// using comma, okay idiom

	if v, ok := m["Bond"]; ok {
		fmt.Println(v, ok)
		delete(m, "Bond")
		fmt.Println(m)
	}

}
