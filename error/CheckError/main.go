package main

import (
	"io/ioutil"
)

func main() {
	_, err := ioutil.ReadDir("/")
	CheckError(err)
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
