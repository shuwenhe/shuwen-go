// 计算1-100的阶乘，存入map，用goroutine
package main

import (
	"fmt"
	"time"
)

var (
	myMap = make(map[int]int, 10)
)

func factorial(n int) {
	res := 1
	for i := 2; i < n; i++ {
		res *= i
	}
	myMap[n-1] = res
}

func main() {
	for i := 2; i < 10; i++ {
		go factorial(i)
	}
	time.Sleep(time.Second)

	for k, v := range myMap {
		fmt.Printf("map[%d] = %d\n", k, v)
	}
}
