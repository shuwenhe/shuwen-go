package main

import (
	"fmt"
)

func main() {
	mySlice := []int{1, 2, 3} // 集合
	mySlice2 := []int{4, 5, 6}
	mySlice3 := append(mySlice, mySlice2...) // 第二个参数mySlice2后面加上3个点...，加上...省略号相当于把mySlice2包含的所有元素打散后传入
	for k, v := range mySlice3 {
		fmt.Printf("k =%d ,v = %d ", k, v)
		fmt.Println()
	}
}
