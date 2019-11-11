package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// playing with go routines. We are using package atomic to fix the race condition below.
func main() {
	var counter uint64
	var wg sync.WaitGroup
	wg.Add(100)

	// Within the loop we have to use the atomic package to print it too, otherwise we get a race
	for i := 0; i < 100; i++ {
		go func() {
			atomic.AddUint64(&counter, 1)
			fmt.Println(atomic.LoadUint64(&counter))
			wg.Done()
		}()

	}
	wg.Wait()
	fmt.Println("end value:", counter)
}
