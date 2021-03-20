package main

import (
	"net/http"
)

func main() {
	Transport()
}

func Transport() {
	http.Transport()
	// type Transport struct {
	// Proxy指定用于针对特定请求返回代理的函数
	// 如果该函数返回一个非空的错误，请求将终止并返回错误
	// 如果Proxy为空或者返回一个空的URL指针，将不会使用代理
	// Proxy func(*Request) (*url.URL, error)
	// Dial指定用于创建TCP连接的dial()函数
	// 如果Dial为空，将默认使用net.Dial()函数
	// Dial func(network, addr string) (net.Conn, error)
	// TLSClientConfig指定用于this.Client的TLS配置
	// 如果为空则使用默认配置
	// TLSClientConfig    *tls.Config
	// DisableKeepAlives  bool // 是否取消长连接，默认false启用长连接
	// DisableCompression bool // 是否取消压缩GZip默认false启用压缩
	// 如果MaxIdleConnsPerHost为非零值，它用于控制每个host所需要
	// 保持的最大空闲连接数。如果该值为空，则使用DefaultMaxIdleConnsPerHost
	// }
}
