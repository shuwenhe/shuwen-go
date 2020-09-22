package main

import (
	"fmt"
)

type Interger int

func main() {
	var a Interger = 1
	a.Add(2)
	fmt.Println("a = ", a)
}

func (a *Interger) Add(b Interger) { // 需要修改对象时必须使用指针
	*a += b
}
