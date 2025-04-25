package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

type A struct {
	i int
	m sync.Mutex
}

func (a *A) Inc(delta int) {
	a.m.Lock()
	defer a.m.Unlock()

	a.i += delta
}

func (a *A) Get() int {
	a.m.Lock()
	defer a.m.Unlock()
	return a.i
}

type B struct {
}

func (b *B) II(c int) {
	for i := 0; i < c; i++ {
		a.Inc(1)
	}
	wg.Done()
}

type C struct {
}

func (c *C) DD(count int) {
	for i := 0; i < count; i++ {
		a.Inc(-1)
	}
	wg.Done()
}

type D struct{}

func (d *D) II(count int) {
	for i := 0; i < count; i++ {
		a.Inc(1)
	}
	wg.Done()
}

type E struct{}

func (e *E) DD(count int) {
	for i := 0; i < count; i++ {
		a.Inc(-1)
	}
	wg.Done()
}

var a A
var b B
var c C
var d D
var e E

var wg sync.WaitGroup

func main() {
	fmt.Println("GOMAXPROCS:", runtime.GOMAXPROCS(0))
	t := time.Now()
	var nCount = 100_000_000
	a.i = 0

	wg.Add(4)
	go b.II(nCount)
	go c.DD(nCount)
	go d.II(nCount)
	go e.DD(nCount)

	wg.Wait()

	fmt.Println("Sum", a.Get())
	fmt.Println("Time", time.Since(t))
}
