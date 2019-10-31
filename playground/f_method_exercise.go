package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
	age   int
}

func (r person) speak() person {
	fmt.Println(r.first, r.last, r.age)
	return r
}

func main() {
	ollie := person{
		first: "Ollie",
		last:  "Young",
		age:   31,
	}
	ollie.speak()
}
