// A look at the comma, ok idiom.
package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	//Putting a value on channel c
	go func() {
		c <- 42
	}()

	// v, ok returns the value of the channel v and a bool value of true if the channel has a value.
	v, ok := <-c
	fmt.Println(v, ok)

	close(c)

	// seeing what v, ok returns after we have closed the channel
	v, ok = <-c
	fmt.Println(v, ok)
}
