package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

}

func Client() {
	// 高级封装
	// 自定义http.Client
	// Client类型的结构
	// type Client struct {
	// http.Client类型包含了4个公开的数据成员
	// Transport RoundTripper // 用于确定http请求的创建机制，必须实现http.RoundTripper接口
	// 如果CheckRedirect不为空，客户端将在跟踪http重定向前调用该函数
	// 两个参数req和via分别为即将发起的请求和已经发起的所有请求，最早的已经发起请求在最前面
	// 如果CheckRedirect返回错误，客户端将直接返回错误，不会再发起请求。
	// 如果CheckRedirect为空，Client将采用一种策略，将在10个连续请求后终止
	// CheckRedirect func(req *Request, via []*Request) error // 定义重定向策略
	// 如果Jar为空，Cookie将不会在请求中发送，并会在响应中被忽略
	// Jar CookieJar // 可在http client中以设置Cookie
	// Jar类型必须实现http.CookieJar接口
	// Timeout time.Duration
	// }
	// 创建自定义的HTTP Client
	client := &http.Client{
		CheckRedirect: redirectPolicyFunc,
	}
	resp, err := client.Get("http://www.xstiku.com")
	fmt.Println(resp)
	checkError(err)
	// func NewRequest(method, url string, body io.Reader) (*Request, error) {
	req, err := http.NewRequest("GET", "http://www.xstiku.com", nil)
	checkError(err)
	req.Header.Add("User-Agent", "Our Custom User-Agent")
	req.Header.Add("If-None-Match", `w/"TheFileEtag"`)
	resp, err := client.Do(req)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "err is %s", err.Error())
	}
	return
}
