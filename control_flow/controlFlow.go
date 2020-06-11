package main

import "fmt"

func main() {
	// loops()
	// loops2()
	// loops3()
	// loopBreakContinue()
	// printAscii()
	// conditionalStatements()
	// conditionalInit()
	conditionalSwitch()
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

func loopBreakContinue() {
	i := 0
	for {
		i++
		if i > 20 {
			break
		}
		if i%2 != 0 {
			continue
		}
		fmt.Println(i)

	}
}

func printAscii() {
	for i := 33; i <= 250; i++ {
		fmt.Printf("%v\t%#x\t%#U\n", i, i, i)
	}

}

func conditionalStatements() {

	if true {
		fmt.Println("001")
	}

	if false {
		fmt.Println("002")
	}

	if !true {
		fmt.Println("003")
	}

	if !false {
		fmt.Println("004")
	}

	if 5 == 5 {
		fmt.Println("005")
	}

	if 5 != 5 {
		fmt.Println("006")
	}

}

func conditionalInit() {

	// having multiple statements in single line
	// also limits the scope of variable to the block
	if x := 20; x == 20 {
		fmt.Println("001")
	}
	// fmt.Println("here is next statement")

	x := 10
	if x == 11 {
		fmt.Println("True")
	} else if x%2 == 0 {
		fmt.Println("not else if")
	} else {
		fmt.Println("this is else")
	}

}

func conditionalSwitch() {

	switch {
	case false:
		fmt.Println("This cannot happen")
	case true:
		fmt.Println("This should happen")
		fallthrough
	case (1 == 1):
		fmt.Println(" 1. after fallen through")
		fallthrough
	case (2 != 2):
		fmt.Println(" 2. after fallen through")
	case (3 == 3):
		fmt.Println(" 3. after fallen through")
		fallthrough
	case (3 > 4):
		fmt.Println(" 4. after fallen through")

	default:
		fmt.Println("this is default")
	}
}
