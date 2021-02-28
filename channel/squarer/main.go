package main

import (
	"fmt"
	"time"
)

// 启动一个goroutine生成100个数发送到ch1中
func counter(out chan<- int) { // 单向通道
	for i := 0; i < 10; i++ {
		out <- i
	}
	close(out)
}

// 启动一个goroutine从ch1中取数计算平方数放到ch2中
func squarer(out chan<- int, in <-chan int) {
	for v := range in {
		out <- v * v
	}
	close(out)
}

func printer(in <-chan int) {
	for v := range in {
		fmt.Println(v)
	}
}

// 从ch2中取出平方数并打印
func main() {
	naturals := make(chan int)
	squares := make(chan int)
	go counter(naturals)
	go squarer(squares, naturals)
	printer(squares)
	time.Tick(time.Second)
}
