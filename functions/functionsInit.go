package main

import (
	"fmt"
)

func main() {
	// n := functionInit("Jason")
	// fmt.Println("This is from main: ", n)
	name, ok := multiReturn("Jason", "Bourne")
	fmt.Println(name, ok)
}

// function SYNTAX
// func (r receiver) identifier(parameters) (return(s)) {...}
// PARAMETER - definition of functions
// ARGUMENT - passed when calling the functions
// EVERYTHING in Go is PASS BY VALUE
func functionInit(s string) string {
	fmt.Println("Hello, ", s)
	return s
}

func multiReturn(first string, last string) (string, bool) {
	fmt.Println(first, last)
	return last, true
}
