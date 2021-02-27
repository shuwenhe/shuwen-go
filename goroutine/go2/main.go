package main

import (
	"fmt"
	"time"
)

func main() {
	go2()
}

func go2() {
	count := 1000
	for i := 0; i < count; i++ {
		go func(i int) {
			for {
				fmt.Printf("hello,shuwen %d\n", i)
			}
		}(i)
		time.Sleep(time.Millisecond)
	}
}
