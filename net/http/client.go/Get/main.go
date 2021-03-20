package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() { // 请求一个网站首页
	resp, err := http.Get("http://www.xstiku.com/user/index.php") // 学思题库首页
	checkError(err)
	fmt.Println("resp = ", resp)
	defer resp.Body.Close()
	io.Copy(os.Stdout, resp.Body) // 将其网页内容打印到标准输出流中
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
	}
}
