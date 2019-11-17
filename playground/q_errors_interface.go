// Here is a program demonstrating use of the builtin type error interface

package main

import (
	"fmt"
)

type customErr struct {
	info string
}

// package builtin contains an error interface that uses the Error() method.
func (r customErr) Error() string {
	return fmt.Sprintf("This is the custom error: %v", r.info)
}

func foo(v error) {
	fmt.Println("foo ran. ", v)
	return
}

func main() {
	myerror := customErr{"I'm tired"}
	foo(myerror)
	// The above is proof that myerror implements the error interface,
	// foo() wouldn't accept it as an argument otherwise.
	// This is because the Error() method for customErr on line 14 implements the error interface.

	fmt.Println("test")
	fmt.Printf("%T\n", myerror)
}
