package main

import (
	"log"
	"os"
)

const (
	UPLOAD_DIR   = "./uploads"
	TEMPLATE_DIR = "./views"
)

func main() {
	Dir()
}

func Dir() {
	f, err := os.Create(UPLOAD_DIR)
	if err != nil {
		panic(err)
	}
	log.Println("f = ", f)
}
