package main

import (
	"fmt"
)

func main() {
	str := "hello"
	for i := 0; i < len(str); i++ { // 0,1,2,3,4
		ch := str[i] // byte即uint8 0~255 2^10=1024 2^9=512 2^8=256 2^7=128
		fmt.Println(i, ch)
	}
}
