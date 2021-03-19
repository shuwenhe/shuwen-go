package main

import (
	"fmt"
	"net/http"
)

func main() {
	files := http.FileServer(http.Dir("/public"))
	fmt.Println("files = ", files)
}
