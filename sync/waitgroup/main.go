package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	wg   sync.WaitGroup
	exit bool
)

func main() {
	wg.Add(3)
	go worker()
	time.Sleep(time.Second * 5) // 业务逻辑执行时间
	exit = true
	wg.Wait()
	fmt.Println("finish...")
}

func worker() {
	defer wg.Done() // 延迟关闭
	for {
		fmt.Println("学思题库-书文教育")
		time.Sleep(time.Second)
		if exit {
			break
		}
	}
}
