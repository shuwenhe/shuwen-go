package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(1)
	go worker(ctx)
	time.Sleep(time.Second * 5)
	cancel()
	wg.Wait()
	fmt.Println("finish...")
}

func worker(ctx context.Context) {
	defer wg.Done()
LABEL:
	for {
		fmt.Println("worker...")
		time.Sleep(time.Second)
		select {
		case <-ctx.Done():
			break LABEL
		default:
		}
	}
}
