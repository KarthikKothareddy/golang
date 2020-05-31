package main

import "fmt"

func main() {
	exercise2()
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
