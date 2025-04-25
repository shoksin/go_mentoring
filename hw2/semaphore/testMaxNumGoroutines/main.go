package main

import (
	"fmt"
	"runtime"
	"time"
)

var x = 0

func main() {
	for i := 0; i < 10_000_000; i++ {
		go doWork()
		if i%10_000 == 0 {
			fmt.Println("goNum =", runtime.NumGoroutine())
		}
	}

	fmt.Println(x)
}

func doWork() {
	x++
	time.Sleep(time.Millisecond * 1)
}
