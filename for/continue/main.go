package main

import (
	"log"
)

func main() {
	continueFor()
}

func continueFor() {
	for i := 0; i < 10; i++ {
		if i == 5 {
			continue // 继续下一次循环跳过i=5的本次循环
		}
		log.Println("i = ", i)
	}
	log.Println("finish")
}
