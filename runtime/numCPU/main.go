package main

import (
	"fmt"
	"runtime"
)

func main() {
	numCPU()
}

func numCPU() {
	fmt.Println("numCPU = ", runtime.NumCPU())
	fmt.Println(runtime.GOMAXPROCS(runtime.NumCPU() - 1))
}
