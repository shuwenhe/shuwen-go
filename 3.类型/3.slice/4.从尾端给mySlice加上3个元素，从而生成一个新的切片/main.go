package main

import (
	"fmt"
)

func main() {
	mySlice := make([]int, 5, 10)
	mySlice2 := append(mySlice, 1, 2, 3)
	for k, v := range mySlice2 {
		fmt.Printf("k =%d ,v = %d ", k, v)
		fmt.Println()
	}
}
