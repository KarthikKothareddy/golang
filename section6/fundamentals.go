package main

import "fmt"

// Efficient Execution
// Efficient Compilation
// Ease of Programming

func main() {
	fmt.Println("This is Programming fundamentals section")
	// varBoolean()
	numericTypes()
}

func varBoolean() {

	a := 7
	b := 42
	fmt.Println(a == b, a >= b, a < b)
}

var x int
var y float64

func numericTypes() {

	// integers
	x = 10
	// x = 1.234 -- fails
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	// floating points
	y = 10.12
	fmt.Println(y)
	fmt.Printf("%T\n", y)
}
