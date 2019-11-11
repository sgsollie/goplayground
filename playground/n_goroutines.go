package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	for i := 0; i <= 100; i++ {
		wg.Add(1)
		go func() {
			fmt.Println("Goroutine", i)
			fmt.Println(runtime.NumGoroutine())
			wg.Done()

		}()
	}
	wg.Wait()
}
