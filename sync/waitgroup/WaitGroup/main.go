package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(3)
	for i := 0; i < 3; i++ {
		go hello(i)
	}
	wg.Wait()
}

func hello(i int) {
	defer wg.Done()
	fmt.Printf("hello%d\n", i)
}
