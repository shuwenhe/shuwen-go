package main

import (
	"fmt"
	"sync"
)

var (
	a  []int
	ch chan int
	wg sync.WaitGroup
)

func main() {
	ch = make(chan int) // a <=> b
	wg.Add(1)           // counter
	go func() {
		defer wg.Done() // defer close
		i := <-ch       // block
		fmt.Println("i = ", i)
	}()
	ch <- 10
	wg.Wait() // Tell main goroutine close
}
