package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	// func HandleFunc(pattern string, handler func(ResponseWriter, *Request)) {
	http.HandleFunc("/hello", helloHanlder)
	err := http.ListenAndServe(":8080", nil)
	checkErr(err)
}

// r is client
func helloHanlder(w http.ResponseWriter, r *http.Request) {
	// func WriteString(w Writer, s string) (n int, err error) {
	// hello string write to w
	io.WriteString(w, "Hello,world!")
}

func checkErr(err error) {
	if err != nil {
		log.Fatal("ListenAndServe:", err.Error())
	}
	return
}
