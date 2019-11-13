package main

import (
	"fmt"
)

func main() {
	c := gen()
	receive(c)

	fmt.Println("about to exit")
}

// func gen() generate numbers on channel and returns c as a receieve from channel.
//CHANNELS BLOCK! So as this function (within its own goroutine) SENDS a value to channel c <-
//.. another function has to pick up (RECEIVE) the values over each iteraton in the loop,
// this is what the receive() function is doing. So as each iteration happens:
// receive picks up the val, prints it, which unblocks the chan,
// allowing the iteration to continue until i reaches 100
// when i reaches 100, the channel is closed with close().

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
