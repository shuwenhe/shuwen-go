package main

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
)

func main() {
	PostForm()
}

func PostForm() {
	// func PostForm(url string, data url.Values) (resp *Response, err error) {
	// The Content-Type header is set to application/x-www-form-urlencoded.
	// http.PostForm()方法实现了标准编码格式为application/x-www-form-urlencoded的表单提交
	// 模拟http表单提交一篇新文章
	resp, err := http.PostForm("http://www.xstiku.com/posts", url.Values{
		"title":   {"article title"},
		"content": {"article body"}},
	)
	fmt.Println(resp)
	checkErr(err)
}

func checkErr(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "err is %s", err.Error())
	}
	return
}
