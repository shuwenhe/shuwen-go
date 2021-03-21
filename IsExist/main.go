package main

import (
	"fmt"
	"os"
)

const (
	DIR = "/"
)

func main() {
	b := IsExist(DIR)
	fmt.Println("b = ", b)
}

func IsExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		os.IsExist(err)
	}
	return true
}
