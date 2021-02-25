package main

import (
	"fmt"
)

func main() {
	i := 2
	switch i { // 左花括号{必须与switch处于同一行
	case 0:
		fmt.Println("0")
	case 1:
		fmt.Println("1")
	case 2:
		fallthrough // 继续执行紧跟的下一个case
	case 3:
		fmt.Println("3")
	case 4, 5, 6: // 单个case中，可以出现多个结果选项
		fmt.Println("4,5,6")
	default:
		fmt.Println("Default")
	}
}
