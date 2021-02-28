package main

import (
	"fmt"
	"time"
)

func main() {
	max := 100
	for i := 0; i < max; i++ {
		go func(i int) {
			fmt.Println("hello", i)
		}(i)
	}
	fmt.Println("main")
	time.Sleep(time.Second)
}
