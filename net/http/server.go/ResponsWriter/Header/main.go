package main

import (
	"net/http"
)

type Post struct {
	User    string
	Threads []string
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/header", header)
	server.ListenAndServe()
}

func header(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Location", "http://xstiku.com")
	w.WriteHeader(302)
}
