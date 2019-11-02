package main

import (
	"fmt"
)

var z int = 0

//foo returns memory address of x (& operator)
func foo() *int {
	x := 0
	return &x
}

func main() {
	y := foo()
	fmt.Println(y)
	*y = 1		// since y = foo (ie, the memory address of x) we can change the value at that address 
			// by dereferencing with * then assigning a val
	fmt.Println(*y)
	
	
	fmt.Println(&z)
	a := &z // setting a to the memory address of z.
	
	fmt.Println(*a)
	*a = 30  // setting the value at a's (and z's) memory address
	fmt.Println(z) // so therefore z should share the be the value of *a (dereferenced a)
	
}
