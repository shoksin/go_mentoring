package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var value int32 = 100
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		swapped := atomic.CompareAndSwapInt32(&value, 100, 200)
		fmt.Println("Swapped 100 -> 200:", swapped)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		swapped := atomic.CompareAndSwapInt32(&value, 100, 300)
		fmt.Println("Swapped 100 -> 300:", swapped)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		swapped := atomic.CompareAndSwapInt32(&value, 100, 400)
		fmt.Println("Swapped 100 -> 400:", swapped)
	}()

	wg.Wait()
	fmt.Println("Final value:", value) // Will be 200 or 300 or 400, depending on swap success
}
