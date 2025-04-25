package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var value int32 = 0
	var wg sync.WaitGroup

	for i := 1; i <= 1000000; i++ {
		wg.Add(1)
		go func(v int32) {
			defer wg.Done()
			atomic.StoreInt32(&value, v) // Safely store a new value
			fmt.Println("Stored value:", v)
		}(int32(i))
	}

	wg.Wait()
	fmt.Println("Final stored value:", atomic.LoadInt32(&value)) // The last stored value
}
