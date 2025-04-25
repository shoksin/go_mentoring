package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int, 2)
	c <- 1
	c <- 2
	go func() {
		time.Sleep(1 * time.Second)
		close(c)
	}()
	time.Sleep(time.Second * 2)
	x := <-c
	fmt.Println("x =", x)
	fmt.Println("done")
}
