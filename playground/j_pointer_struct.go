package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
}

//Function to change the first name of a person, taking pointer to type person as an argument
func changeMe(p *person) {
	(*p).first = "Blank" 	// change first name to "Blank"

}

func main() {
	p1 := person{
		first: "Ollie",
		last:  "Young",
	}
	fmt.Println(p1)

	changeMe(&p1) // takes a pointer to p1.
	fmt.Println(p1)
}

