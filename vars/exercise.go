package main

import "fmt"

func main() {
	exercise5()
}

func exercise1() {

	// short declaration operator
	x := 42
	y := "James Bond"
	z := true

	fmt.Println(x, y, z)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}

var X int = 42
var Y string = "James Bond"
var Z bool = true

func exercise2() {
	// thse will print out 0, , false
	// these are called ZERO VALUES
	// fmt.Println(X)
	// fmt.Println(Y)
	// fmt.Println(Z)

	s := fmt.Sprintf("%v%v%v", X, Y, Z)
	fmt.Println(s)

}

type myType int

var x myType

func exercise4() {
	fmt.Printf("%T\n", x)
	fmt.Println(x)
	x = 42
	fmt.Println(x)
}

var y int

func exercise5() {
	x = 42
	y = int(x)
	fmt.Printf("%T\n", y)
	fmt.Println(y)

}
