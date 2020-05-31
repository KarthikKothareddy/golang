package main

import "fmt"

func main() {
	//fmt.Println("=")
	fmt.Println("Hello, world!")

	// control flow
	// 1. sequence
	// 2. loop; iterative
	// 3. conditional

	// foo
	// bar()
	typeTesting()
}

func foo() {
	for i := 0; i < 5; i++ {
		if i%2 == 0 {
			n, _ := fmt.Println(i, " ")
			fmt.Println(n)
			fmt.Println("Hey I'm learning GoLang")
		}
	}

}

func bar() {

	// declaration and assignment
	// := short declaration operator
	x := 5

	fmt.Println(x)
	x = 10 + 2
	fmt.Println(x)

	// declarative
	var y = 22
	fmt.Println(y)
}

// this is a static Programming language, not dynamic
var y int = 99
var z string = `
Hi I am 
"Karthik"
`

var z1 string
var y1 int

var x = 11

func typeTesting() {

	fmt.Println(x)
	// type
	fmt.Printf("%T\n", x)
	// binary
	fmt.Printf("%b\n", x)
	// hexadecimal
	fmt.Printf("%x\n", x)
	// hexadecimal with '0x' infront of it
	fmt.Printf("%#x\n", x)
}
