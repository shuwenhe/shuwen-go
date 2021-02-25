package main

import (
	"fmt"
)

func main() {
	var myArray [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} // 先定义一个数组
	// mySlice := myArray[:5]                                       // array & slcie 集合论 《离散数学》屈婉玲
	var mySlice []int = myArray[:5] // 基于数组创建一个数组切片，array & slcie 集合论 《离散数学》屈婉玲

	fmt.Println("myArray = ")
	for _, vArray := range myArray {
		fmt.Println(vArray)
	}

	fmt.Println("mySlice = ")
	for _, vSlice := range mySlice {
		fmt.Println(vSlice)
	}
}
