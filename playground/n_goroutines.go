package main

import (
	"fmt"
	"sync"
)

// playing with go routines. We are using a mutex to lock the counter variable below. This prevents a race condition
func main() {
	var mu sync.Mutex
	var wg sync.WaitGroup
	counter := 0
	wg.Add(100)

	for i := 0; i < 100; i++ {
		go func() {
			mu.Lock()
			counter++
			fmt.Println("Counter: ", counter)
			mu.Unlock()
			wg.Done()
		}()

	}
	wg.Wait()
	fmt.Println("end value:", counter)
}
