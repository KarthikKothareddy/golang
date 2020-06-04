package main

import (
	"fmt"
	"runtime"
)

// Efficient Execution
// Efficient Compilation
// Ease of Programming

func main() {
	fmt.Println("This is Programming fundamentals section")

	// runtime package
	fmt.Println("Go OS: ", runtime.GOOS)
	fmt.Println("GO Architecture: ", runtime.GOARCH)

	// varBoolean()
	// numericTypes()
	stringTypes()
}

func varBoolean() {

	a := 7
	b := 42
	fmt.Println(a == b, a >= b, a < b)
}

func numericTypes() {

	var x int
	var y float64

	// integers
	x = 10
	// x = 1.234 -- fails
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	// floating points
	y = 10.12
	fmt.Println(y)
	fmt.Printf("%T\n", y)

	var i int8 = -128
	fmt.Println(i)

}

func stringTypes() {

	s := "Hello, World!"
	fmt.Println(s)
	fmt.Printf("%T\n", s)

	// conversion from type string to slice of bytes ascii
	bs := []byte(s)
	fmt.Println(bs)
	fmt.Printf("%T\n", bs)

	for i := 0; i < len(s); i++ {
		// utf-8 codepoints
		fmt.Printf("%#U ", s[i])
	}

	fmt.Println("")

	for i, v := range s {
		// hexadecimal
		fmt.Printf("At index position %d we have hex %#x \n", i, v)
	}

}

func numeralSystems() {

	// as humans we use base 10 numeral system
	n := 10
	fmt.Println(n)
}

func constants() {

}
