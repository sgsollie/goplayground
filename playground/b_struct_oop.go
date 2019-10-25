// more playing with structs, with an emphasis on type (basic OOP in go)

package main

import (
	"fmt"
)

type person struct {
	first    string
	last     string
	icecream []string
}

func main() {
	// Setting up my VALUES of TYPE person with SLICE of STRING for favourite icecream flavours.
	p1 := person{
		first:    "Ollie",
		last:     "Young",
		icecream: []string{"Salted Caramel", "Chocolate", "CookieDough"},
	}

	p2 := person{
		first:    "James",
		last:     "Bond",
		icecream: []string{"Strawberry", "Martini", "Bubblegum"},
	}

	// Declaring people as a SLICE of TYPE person
	people := []person{p1, p2}

	//Ranging over each person's favourite icecream
	for i, v := range people {
		fmt.Printf("%v: %v\n", i, v)
		for j, w := range v.icecream {
			fmt.Printf("\t%v: %v\n", j, w)
		}
	}
}
