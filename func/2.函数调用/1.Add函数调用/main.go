package main

import (
	"fmt"
	"shuwen-go/5.函数/2.函数调用/1.Add函数调用/mymath"
)

func main() {
	c, _ := mymath.Add(1, 2)
	fmt.Println("c = ", c)
}
