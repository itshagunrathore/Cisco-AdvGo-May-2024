package main

import (
	"context"
	"fmt"
	"time"
)

// consumer
func main() {
	rootCtx := context.Background()
	cancelCtx, cancel := context.WithCancel(rootCtx)
	fmt.Println("Hit ENTER to stop....")
	go func() {
		fmt.Scanln()
		cancel()
	}()
	ch := genNos(cancelCtx)
	for data := range ch {
		fmt.Println(data)
		time.Sleep(500 * time.Millisecond)
	}
	fmt.Println("Done")
}

// producer
func genNos(ctx context.Context) <-chan int {
	ch := make(chan int)
	go func() {
	LOOP:
		for i := 1; ; i++ {
			select {
			case <-ctx.Done():
				break LOOP
			default:
				ch <- 10 * i
			}
		}
		fmt.Println("cancellation signal received")
		close(ch)
	}()
	return ch
}
