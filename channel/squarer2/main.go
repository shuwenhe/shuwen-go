package main

import (
	"fmt"
)

func main() { // 计算1-9的平方
	ch1 := make(chan int) // 将1-9存入通道ch1
	ch2 := make(chan int) // 将1-9从ch1中读取再计算平方数并存入ch2通道
	go counter(ch1)
	go squarer(ch2, ch1)
	printer(ch2)
}

func counter(ch1 chan<- int) { // 单向通道
	for i := 1; i < 100; i++ {
		ch1 <- i
	}
	close(ch1)
}

func squarer(ch2 chan<- int, ch1 <-chan int) {
	for v := range ch1 {
		ch2 <- v * v
	}
	close(ch2)
}

func printer(ch2 <-chan int) {
	for v := range ch2 {
		fmt.Println(v)
	}
}
