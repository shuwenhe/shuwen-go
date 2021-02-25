package main

import (
	"fmt"
)

func main() {
	str := "hello" // string(字符串) -> "hello" ; char(字节)(UTF-8) -> "h"
	for k, v := range str {
		fmt.Println(k, v) // v的类型为rune,代表UTF-8字符串的单个字节的值
	}
}
