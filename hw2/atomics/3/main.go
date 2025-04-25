package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var value int32 = 100
	var wg sync.WaitGroup

	for i := int32(1); i <= 10; i++ {
		wg.Add(1)
		go func(v int32) {
			defer wg.Done()
			oldValue := atomic.SwapInt32(&value, v)
			fmt.Printf("Swapped value: old = %d, new = %d\n", oldValue, v)
		}(i * 100)
	}

	wg.Wait()
	fmt.Println("Final value:", value)
}
