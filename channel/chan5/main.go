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

func noBufChannel() {
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

func bufChannel() {
	ch = make(chan int, 2) // buf
	ch <- 1
	ch <- 3
	i := <-ch
	fmt.Println("i = ", i)
	j := <-ch
	fmt.Println("j = ", j)
	close(ch) // Close channel
}

func main() {
	bufChannel()
}
