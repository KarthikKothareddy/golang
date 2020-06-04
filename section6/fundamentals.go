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
	// stringTypes()
	// constants()
	// iotaPractice()
	// bitShifting()
	bitShitftingIota()
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

	// const is a keyword
	const a = 11
	const b = 22

	fmt.Println(a)
	fmt.Printf("%T\n", a)
	fmt.Println(b)
	fmt.Printf("%T\n", b)

	const (
		c float64 = 33.33
		d string  = "GoLang"
	)

	fmt.Println(c)
	fmt.Printf("%T\n", c)
	fmt.Println(d)
	fmt.Printf("%T\n", d)

}

func iotaPractice() {

	const (
		a = iota
		b
		c
	)

	const (
		d = iota
		e
		f
	)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
}

func bitShifting() {

	x := 4
	fmt.Printf("%d\t\t%b\n", x, x)

	// shift bits
	y := x << 1
	fmt.Printf("%d\t\t%b", y, y)

}

func bitShitftingIota() {

	kb := 1024
	mb := 1024 * kb
	gb := 1024 * mb

	fmt.Printf("%d\t\t%b\n", kb, kb)
	fmt.Printf("%d\t\t%b\n", mb, mb)
	fmt.Printf("%d\t\t%b\n", gb, gb)

	// same functionality with iota
	const (
		kb1 = 1024
		mb1 = 1024 * kb1
		gb1 = 1024 * mb1
	)

	fmt.Printf("%d\t\t%b\n", kb1, kb1)
	fmt.Printf("%d\t\t%b\n", mb1, mb1)
	fmt.Printf("%d\t\t%b\n", gb1, gb1)

	// iota with bitshifting NINJA LEVEL

	const (
		_   = iota
		kb2 = 1 << (iota * 10)
		mb2 = 1 << (iota * 10)
		gb2 = 1 << (iota * 10)
	)

	fmt.Printf("%d\t\t%b\n", kb2, kb2)
	fmt.Printf("%d\t\t%b\n", mb2, mb2)
	fmt.Printf("%d\t\t%b\n", gb2, gb2)

}
