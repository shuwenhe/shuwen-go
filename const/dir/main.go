package main

import (
	"os"

	"github.com/shuwenhe/shuwen-go/errors"
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
	errors.CheckError()
}
