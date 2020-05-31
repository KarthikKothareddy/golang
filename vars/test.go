package main

import "fmt"

func main() {
	fmt.Println("Hello, world!")

	// control flow
	// 1. sequence
	// 2. loop; iterative
	// 3. conditional
	for i := 0; i < 5; i++ {
		if i%2 == 0 {
			n, _ := fmt.Println(i, " ")
			fmt.Println(n)
			foo()
		}
	}

}

func foo() {
	fmt.Println("Hey I'm learning GoLang")
}
