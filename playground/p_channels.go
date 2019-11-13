package main

import (
	"fmt"
)

func main() {
	c := gen()
	receive(c)

	fmt.Println("about to exit")
}

//generate numbers on channel - returns c as a receieve from channel
//closes c after counting to 100
func gen() <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
	}()
	return c
}

//func that is receiving all data from c in main() scope and printing it out until the channel
//is closed in the other goroutine on line 19.
func receive(ch <-chan int) {
	for v := range ch {
		fmt.Println(v)
	}

}
