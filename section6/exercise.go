package main

import (
	"fmt"
)

func main() {
	// exercise1()
	// exercise2()
	exercise3()
}

func exercise1() {

	x := 911

	fmt.Printf("%d\n", x)
	fmt.Printf("%b\n", x)
	fmt.Printf("%#x\n", x)

}

func exercise2() {

	a := (10 == 10)
	b := (10 <= 5)
	c := (10 >= 1)
	d := (1 != 2)
	e := (10 < 5)
	f := (10 > 5)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)

}

func exercise3() {

	const a int = 2
	fmt.Printf("%d\t%T\n", a, a)

	const b = "Karthik"
	fmt.Printf("%s\t%T\n", b, b)

	const (
		c int = 10
		d     = "Hello"
	)
	fmt.Printf("%d\t", c)
	fmt.Printf("%T\n", c)

}
