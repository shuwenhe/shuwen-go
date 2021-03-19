package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", index) // 将发送到根URL的请求重定向到处理器
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) { // 处理器
	fmt.Fprintf(w, "Hello world")
}
