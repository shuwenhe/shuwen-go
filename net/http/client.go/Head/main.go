package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	Head()
}

func Head() {
	// func Head(url string) (resp *Response, err error) {
	// http中的Head请求方法只请求目标URL的头部信息，即HTTPHeader而不返回HTTP的Body
	resp, err := http.Head("http://www.xstiku.com/")
	fmt.Println(resp)
	checkErr(err)
}

func checkErr(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "err is %s", err.Error())
	}
	return
}
