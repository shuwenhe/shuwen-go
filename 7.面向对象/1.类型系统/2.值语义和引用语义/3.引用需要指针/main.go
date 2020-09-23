package main

import (
	"fmt"
)

func main() {
	var a = [3]int{1, 2, 3}
	var b = &a
	b[1]++
	fmt.Println("a = ", a)  // a =  [1 3 3]
	fmt.Println("b = ", *b) // b =  [1 3 3]
}
