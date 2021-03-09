package main

import (
	"fmt"
	"net/http"
)

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/writeheader", WriteHeader)
	server.ListenAndServe()
}

func WriteHeader(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(501)
	fmt.Fprintln(w, "No such service,try next door.")
}
