package main

import (
	"fmt"
	"net/http"
)

func main() {
	RequestURL()
}

func RequestURL() {
	var r *http.Request
	fmt.Println("r.URL = ", r.URL.Path)
}
