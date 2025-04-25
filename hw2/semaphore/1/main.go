package main

import (
	"fmt"
	"net/http"
	"sync"
)

type Semaphore struct {
	c chan struct{}
}

func NewSemaphore(n int) *Semaphore {
	return &Semaphore{make(chan struct{}, n)}
}

func (s *Semaphore) Acquire() {
	s.c <- struct{}{}
}

func (s *Semaphore) Release() {
	<-s.c
}

var wg = &sync.WaitGroup{}

func main() {
	sem := NewSemaphore(4)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go doJob(i, sem)
	}
	wg.Wait()
	fmt.Println("Work finished")
}

func doJob(i int, sem *Semaphore) {
	sem.Acquire()
	defer wg.Done()
	defer sem.Release()
	resp, err := http.Get("https://google.com/")
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	fmt.Printf("doJob: %d, statusCode: %d response: %s\n", i, resp.StatusCode, resp.Header.Get("content-type"))
}
