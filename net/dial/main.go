package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	dial()
}

func dial() {
	fmt.Println("conn = ")
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Print(err)
	}
	fmt.Println("conn = ", conn)
}
