package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Book struct {
	Title       string
	Authors     []string
	Publisher   string
	IsPublished bool
	Price       float64
}

func main() {
	Marshal()
}

func Marshal() {
	gobook := Book{ // Book类型的实例对象
		"shuwen-os",
		[]string{"Shuwen", "Richard", "Ritchie"}, // 添加[]string
		"xstiku.com",
		true,
		9.9,
	}
	b, err := json.Marshal(gobook) // func Marshal(v interface{}) ([]byte, error) {
	json.Unmarshal
	fmt.Println("b = ", b)         // struct->[]byte
	fmt.Println("b = ", string(b)) // []byte->string
	checkErr(err)
}

func checkErr(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "err is %s", err.Error())
	}
	return
}
