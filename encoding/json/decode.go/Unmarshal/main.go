package main

import (
	"encoding/json"
	"fmt"
)

type Book struct {
	Title       string
	Authors     []string
	Publisher   string
	IsPublished bool
	Price       float64
}

func main() { // 解码JSON数据
	Unmarshal()
}

// Go是如何将JSON数据解码后的值一一准确无误地关联到一个数据结构中的相应字段呢？
func Unmarshal() {
	// func Unmarshal(data []byte, v interface{}) error {
	// data []byte : JSON格式文本(比特序列)
	// v interface{} : 目标输出容器,用于存放解码后的值
	// 要解码一段JSON数据，首先需要在Go中创建一个目标类型的实例对象，用于存放解码后的值
	var book Book
	b := []byte(`{"Title":"shuwen-os","Authors":["Shuwen","Richard","Ritchie"],"Publisher":"Shuwen University","IsPublished":true,"Price":9.9}`)
	// json.Unmarshal()函数会根据一个约定的顺序查找目标结构中的字段，如果找到一个即发生匹配
	json.Unmarshal(b, &book) // 将JSON文本->解码为Go的数据结构 string->struct...
	fmt.Println("book = ", book)
	fmt.Println("title = ", book.Title)
	fmt.Println("publisher = ", book.Publisher)
}
