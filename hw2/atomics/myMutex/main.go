package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type MyMutex struct {
	i int64
}

func (m *MyMutex) set(value int64) {
	atomic.StoreInt64(&m.i, value)
}

func (m *MyMutex) get() int64 {
	return atomic.LoadInt64(&m.i)
}

func (m *MyMutex) Lock() {
	for atomic.CompareAndSwapInt64(&m.i, 0, 1) != true {
	}
	return
}

func (m *MyMutex) Unlock() {
	if m.get() == 0 {
		panic(fmt.Errorf("double unlocking"))
	}
	m.set(0)
}

func main() {
	var wg sync.WaitGroup
	m := &MyMutex{}
	count := 0
	for i := 0; i < 100000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			m.Lock()
			count++
			m.Unlock()
		}()
	}
	wg.Wait()
	fmt.Println(count)
}
