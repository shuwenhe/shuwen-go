package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	ReadDir()
}

func ReadDir() {
	fileInfoAll, err := ioutil.ReadDir("/Users/shuwen/go/src/shuwen/shuwen-go/io/ioutil/ioutil.go/ReadDir")
	if err != nil {
		fmt.Println("err = ", err)
	}
	for _, fileInfo := range fileInfoAll {
		fmt.Println("fileInfo = \n", fileInfo)
	}
}
