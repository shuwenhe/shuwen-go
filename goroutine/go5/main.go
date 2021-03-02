package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go test(i)
	}
	for j := 0; j < 10; j++ {
		fmt.Printf("j = %d\n", j)
	}
	wg.Wait()
}

func test(i int) {
	defer wg.Done()
	fmt.Printf("hello%d\n", i)
}
