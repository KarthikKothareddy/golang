package main

import (
	"fmt"
)

func main() {
	loops()
}

func loops() {
	// there is no while loop in Go
	// for init; condition; post{}

	for i := 0; i <= 100; i++ {
		fmt.Println("Hello, Golang")
	}

}
