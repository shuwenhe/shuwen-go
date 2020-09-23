package main

import "fmt"

func main() { // 当main返回时，程序就退出，程序不等待其他goroutine(非主goroutine)
	for i := 0; i < 10; i++ {
		// Add(i, i)
		go Add(i, i)
	}
}

// Add 两个参数相加
func Add(x, y int) {
	z := x + y
	fmt.Println("z = ", z)
}
