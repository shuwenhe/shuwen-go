package main

import (
	"fmt"
)

func main() {
	var v1 int = 1
	var v2 int64 = 123
	var v3 string = "hello"
	var v4 float32 = 1.23
	MyPrintf(v1, v2, v3, v4)
}

// MyPrintf 输出interface
func MyPrintf(args ...interface{}) {
	for _, arg := range args {
		switch arg.(type) {
		case int:
			fmt.Println(arg, "is an int value.")
		case int64:
			fmt.Println(arg, "is an int64 value.")
		case string:
			fmt.Println(arg, "is a string value")
		default:
			fmt.Println(arg, "is an unknown type.")
		}
	}
}
