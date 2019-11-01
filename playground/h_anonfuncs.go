package main

import (
	"fmt"
)

func main() {
	x := func() {
		fmt.Println("This is an anonymous func assigned to a variable")
	}
	x()

	// Running an anonymous func directly
	func() {
		fmt.Println("Ths is an anon func to print decimal numbers 0 to 10 in binary:")
		for z := 0; z <= 10; z++ {
			fmt.Printf("%b\t", z)
		}
	}()
}
