package main

import (
	"fmt"
	"log"
)

func main() {
	breakFor()
}

func breakFor() {
	for i := 0; i < 10; i++ {
		if i == 7 {
			break
		}
		log.Println("i = ", i)
	}
	fmt.Println("finish")
}
