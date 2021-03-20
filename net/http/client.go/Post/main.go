package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

}

func post() {
	// func Post(url, contentType string, body io.Reader) (resp *Response, err error) {
	// url->请求的URL地址
	// contentType->POST资源类型
	// body->比特流io
	resp, err := http.Post("http://www.xstiku.com/upload", "image/jpeg", &imageDateBuf)
	checkErr(err)

	if resp.StatusCode != http.StatusOK {
		return
	}
}

func checkErr(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "err: %s", err.Error())
	}
	return
}
