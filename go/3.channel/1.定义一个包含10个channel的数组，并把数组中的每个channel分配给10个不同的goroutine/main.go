package main

import (
	"fmt"
)

func main() {
	chs := make([]chan int, 10) // 定义10个channel数组
	for i := 0; i < 10; i++ {
		chs[i] = make(chan int)
		go Count(chs[i]) // 并发执行Count
	}
	for _, ch := range chs {
		<-ch
	}
}

// Count 计数
func Count(ch chan int) {
	fmt.Println("Counting")
	ch <- 1
}
