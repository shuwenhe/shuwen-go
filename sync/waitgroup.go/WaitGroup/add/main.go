package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go Add(i, i)
	}
	wg.Wait()
}

func Add(x, y int) {
	wg.Done()
	fmt.Println(x + y)
}
