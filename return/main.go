package main

import (
	"fmt"
)

func example(x int) int {
	if x == 0 {
		return 5
	} else {
		return x
	}
}

func main() {
	x := example(2)
	fmt.Println("x = ", x)
}
