package main

import (
	"fmt"
)

func main() {
	myFunc(1, 2, 3)
}

func myFunc(args ...int) {
	for _, arg := range args {
		fmt.Println("arg = ", arg)
	}
}
