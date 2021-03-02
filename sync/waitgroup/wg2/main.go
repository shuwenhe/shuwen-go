package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(3)
	for i := 0; i < 3; i++ {
		go Add(i, i)
	}
	wg.Wait()
}

func Add(x, y int) {
	defer wg.Done()
	z := x + y
	fmt.Println(z)
}
