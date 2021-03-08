package main

import (
	"net/http"
)

func main() {
	ListenAndServeTLS()
}

func ListenAndServeTLS() { // SSL=>TLS(Transport Layer Security传输层安全协议)
	server := http.Server{ // https => 在TLS连接的上层进行http通信
		Addr:    "127.0.0.1:8080", // SSL更常用
		Handler: nil,
	}
	server.ListenAndServeTLS("cert.pem", "key.pem") // cert.pem文件SSL证书 key.pem服务器私钥
}
