package main

import (
	"fmt"
	"math/rand"
	"time"
)

// consumer
func main() {
	ch := genNos()
	for data := range ch {
		fmt.Println(data)
		time.Sleep(500 * time.Millisecond)
	}
}

// producer
func genNos() <-chan int {
	ch := make(chan int)
	go func() {
		count := rand.Intn(20)
		fmt.Printf("[@genNos], count = %d\n", count)
		for i := 1; i <= count; i++ {
			ch <- 10 * i
		}
		close(ch)
	}()
	return ch
}
