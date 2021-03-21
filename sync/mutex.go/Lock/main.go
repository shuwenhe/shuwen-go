package main

import (
	"log"
	"runtime"
	"sync"
)

var counter int = 0 //10个goroutine共享变量counter

func main() {
	Mutex()
}

func Mutex() {
	lock := &sync.Mutex{}
	for i := 0; i < 10; i++ { // 10个goroutine
		go Count(lock)
	}

	for { // 使用for不断检查counter值，同样需要加锁
		lock.Lock()
		c := counter
		lock.Unlock()
		runtime.Gosched()
		if c >= 10 {
			break
		}
	}
}

func Count(lock *sync.Mutex) {
	lock.Lock()
	counter++
	log.Println("counter : ", counter)
	lock.Unlock()
}
