package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fn()
}

func fn() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 5; i++ {
		v1 := rand.Int()
		v2 := rand.Intn(10) // [0,n) => (0,1,2,3,4,5,6,7,8,9)
		fmt.Printf("v1 =%d ,v2 = %d \n", v1, v2)
	}
}
