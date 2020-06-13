package main

import (
	"fmt"
)

func main() {

	// structureFundamentals()
	// embeddedStructures()
	anonymousStructures()
}

type person struct {
	first string
	last  string
	age   int
}

func structureFundamentals() {

	p1 := person{
		first: "Jason",
		last:  "Bourne",
		age:   32,
	}

	p2 := person{
		first: "Iron",
		last:  "Man",
	}

	fmt.Println(p1, p2)
	fmt.Println(p1.first, p2.last)

}

type secretAgent struct {
	// here we're embedding 'person'
	person
	ltk bool
}

func embeddedStructures() {

	sa1 := secretAgent{
		person: person{
			first: "Jason",
			last:  "Bourne",
			age:   32,
		},
		ltk: true,
	}

	// referencing different level fields
	fmt.Println(sa1.first, sa1.ltk)
	fmt.Println(sa1.person, sa1.person.age)

}

func anonymousStructures() {

	// anonymous because it doesn't have a name and cannot be reused
	p1 := struct {
		first string
		last  string
		age   int
	}{
		first: "Jason",
		last:  "Bourne",
		age:   32,
	}

	fmt.Println(p1)

}
