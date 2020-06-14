package main

import (
	"fmt"
)

func main() {
	functionInit("Jason")
}

// function SYNTAX
// func (r receiver) identifier(parameters) (return(s)) {...}
// PARAMETER - definition of functions
// ARGUMENT - passed when calling the functions
// EVERYTHING in Go is PASS BY VALUE
func functionInit(s string) {
	fmt.Println("Hello, ", s)
}
