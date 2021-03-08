package main

import (
	"net/http"
)

func main() {
	LisenAndServe()
}

func LisenAndServe() {
	http.ListenAndServe("", nil)
}
