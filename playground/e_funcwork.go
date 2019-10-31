package main

import (
	"fmt"
)

func foo(n ...int) int {
	total := 0
	for _, v := range n {
		total += v
	}
	return total

}

func bar(n []int) int {
	total := 0
	for _, v := range n {
		total += v
	}
	return total

}

func main() {
	x := foo(5, 5, 10)
	fmt.Println(x)

	y := bar([]int{5, 5, 5})
	fmt.Println(y)

	//unfurling the slice - since foo takes a variadic parameter of type int, it doesn't take a type of slice of int.
	//so we can pass xx as a slice of int, suffixed with ... which passes each value in the slice as an individual int.
	xx := []int{2, 4, 2, 2}
	fmt.Println(foo(xx...))
}

