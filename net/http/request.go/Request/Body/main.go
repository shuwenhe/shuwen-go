package main

import (
	"fmt"
	"net/http"
)

func body(w http.ResponseWriter, r *http.Request) {
	len := r.ContentLength    // 获取主体字节长度
	body := make([]byte, len) // 根据长度创建字节数组
	r.Body.Read(body)         // 调用Read方法将主体数据读取到字节数组中
	fmt.Fprintln(w, string(body))
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080", // curl -id "first_name=Shuwen&last_name=He" 127.0.0.1:8080/body
	}
	http.HandleFunc("/body", body)
	server.ListenAndServe()
}
