package main

import (
	"fmt"
)

//Put i on channel c with each iteration. - Iterate 10 times.
func gen(c chan int) {
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
	}()
}

//Keep looping over channel c for a value (we know it will receive 100 ints hence the loop)
func receive(c <-chan int) {
	for j := 0; j < 100; j++ {
		fmt.Println(<-c)
	}
}

func main() {
	c := make(chan int)
	// Creating 10 go routines for this exercise.
	// Each go routine is putting 10 values on a channel
	// So 10 vals * 10 Go routines = 100 vals on the channel which need to be recieved.
	for i := 0; i < 10; i++ {
		go gen(c)
	}
	receive(c)
}
