package main

import (
	"log"
)

func main() {
	switchCase()
}

func switchCase() {
	i := 3
	switch i {
	case 1:
		log.Println("Shuwen")
	case 2:
		log.Println("Richard")
	case 3:
		log.Println("Ritchie")
	default:
		log.Println("Sophie")
	}
}
