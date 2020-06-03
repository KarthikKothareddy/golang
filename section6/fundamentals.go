package main

import "fmt"

func main() {
	fmt.Println("This is Programming fundamentals section")
	varBoolean()
}

var x bool

func varBoolean() {

	a := 7
	b := 42
	fmt.Println(a == b, a >= b, a < b)
}
