package main

import (
	"net/http"
)

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/write", writeExample)
	server.ListenAndServe()
}

func writeExample(w http.ResponseWriter, r *http.Request) {
	str := `<html>
	<head><title>Go web </title></head>
	<body><h1>hello world</h1></body>
	</html>`
	w.Write([]byte(str))
}
