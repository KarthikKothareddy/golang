package main

import "fmt"

func main() {
	// loops()
	// loops2()
	loops3()
}

func loops() {
	// there is no while loop in Go
	// for init; condition; post{}

	for i := 0; i <= 3; i++ {
		// fmt.Println("Hello, Golang")
		for j := 0; j <= i; j++ {
			fmt.Println(i, j)
		}
	}

}

func loops2() {
	// for statement with a single condition
	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}

}

func loops3() {
	i := 1
	for {
		if i > 10 {
			break
		}
		fmt.Println(i)
		i++
	}

}
