package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

type Semaphore struct {
	capacity int32
	size     int32
	signalCh chan struct{}
}

func NewSemaphore(n int32) *Semaphore {
	return &Semaphore{
		capacity: n,
		signalCh: make(chan struct{}, n),
	}
}

func (s *Semaphore) Acquire() {
	for {
		size := atomic.LoadInt32(&s.size)
		if size <= s.capacity {
			if atomic.CompareAndSwapInt32(&s.size, size, size+1) {
				s.signalCh <- struct{}{}
				return
			}
		}
	}
}

func (s *Semaphore) Release() {
	if atomic.AddInt32(&s.size, -1) >= 0 {
		<-s.signalCh
	}
}

func main() {
	s := NewSemaphore(10)
	wg := sync.WaitGroup{}
	t := time.Now()
	for i := 0; i < 100; i++ {
		wg.Add(1)
		s.Acquire()
		go func() {
			defer wg.Done()
			defer s.Release()
			fmt.Println(i, "start")
			time.Sleep(time.Second)
			fmt.Println(i, "end")
		}()
	}

	wg.Wait()
	fmt.Println("Final value:", s.size)
	fmt.Println("Time taken:", time.Since(t).Seconds())
}
