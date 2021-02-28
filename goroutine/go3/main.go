package main

import (
	"fmt"
	"time"
)

func main() {
	max := 1048575
	for i := 0; i < max; i++ {
		go hello(i)
	}
	fmt.Println("main")
	time.Sleep(time.Second)
}

func hello(i int) {
	fmt.Println("hello", i)
}
