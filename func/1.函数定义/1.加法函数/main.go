package main

import (
	"errors"
	"fmt"
)

func main() {
	ret, _ := Add(1, 2)
	fmt.Println("ret = ", ret)
}

// Add 加法函数
func Add(a, b int) (ret int, err error) { // a,b相邻参数类型相同，可以合并
	if a < 0 || b < 0 {
		err = errors.New("Should be non-negative numbers!")
		return
	}
	return a + b, nil // 支持多重返回值
}
