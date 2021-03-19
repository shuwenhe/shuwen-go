package main

import (
	"fmt"
	"net/http"
)

func main() { // 通过一些代码将接收到的请求重定向到处理器
	mux := http.NewServeMux() // 创建一个多路复用器
	fmt.Println("mux = ", mux)
}
