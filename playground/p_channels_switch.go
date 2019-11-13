// More work with channels. This is to print out values from two channels using a select statement

package main

import (
	"fmt"
)

func main() {

	q := make(chan int)
	// The below two lines happen concurrently
	c := gen(q)
	receive(c, q)

	fmt.Println("about to exit")
}

func receive(c1, c2 <-chan int) {
	// loop over channel c1 to consume values (placed here by the loop in gen())
	// when channel c1 is finished counting to 100 (line 40) it adds 1 to c2 - when we receive that, we quit.
	for {
		select {
		case msg1 := <-c1:
			fmt.Println(msg1)
		case <-c2:
			fmt.Println("quit")
			return
		}
	}

}

func gen(q chan<- int) <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		q <- 1
		close(c)

	}()
	return c
}
