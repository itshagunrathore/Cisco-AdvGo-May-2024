package main

import (
	"context"
	"fmt"
	"time"
)

type ContextKey struct {
}

// consumer
func main() {
	contextKey := ContextKey{}
	rootCtx := context.Background()
	valCtx := context.WithValue(rootCtx, contextKey, 5)
	timeoutCtx, cancel := context.WithTimeout(valCtx, 5*time.Second)
	fmt.Println("Will timeout after 5 secs.. Hit ENTER to stop manually!")
	go func() {
		fmt.Scanln()
		cancel()
	}()

	ch := genNos(timeoutCtx)
	for data := range ch {
		fmt.Println(data)
		time.Sleep(500 * time.Millisecond)
	}
	fmt.Println("Done")
}

// producer
func genNos(ctx context.Context) <-chan int {
	ch := make(chan int)
	contextKey := ContextKey{}
	multiplier := ctx.Value(contextKey).(int)
	go func() {
	LOOP:
		for i := 1; ; i++ {
			select {
			case <-ctx.Done():
				break LOOP
			default:
				ch <- multiplier * i
			}
		}
		fmt.Println("cancellation signal received")
		fmt.Println("error :", ctx.Err())
		close(ch)
	}()
	return ch
}
