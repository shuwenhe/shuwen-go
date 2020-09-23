package main

import (
	"fmt"
)

func main() {
	var a = [3]int{1, 2, 3} // array是初始化确定大小
	fmt.Println("a = ", a)  // 输出结果：a =  [1 2 3]
	var b = a
	fmt.Println("b = ", b) // 输出结果：b =  [1 2 3]
	b[1]++
	fmt.Println("b = ", b) // 输出结果：b =  [1 3 3]，表明b=a赋值语句是array内容的完整复制
}
