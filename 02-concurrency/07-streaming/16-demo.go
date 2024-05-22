package main

import (
	"fmt"
	"math/rand"
	"time"
)

// consumer
func main() {
	ch := make(chan int)
	go genNos(ch)
	for {
		if data, isOpen := <-ch; isOpen {
			fmt.Println(data)
			time.Sleep(500 * time.Millisecond)
			continue
		}
		break
	}
}

// producer
func genNos(ch chan<- int) {
	count := rand.Intn(20)
	fmt.Printf("[@genNos], count = %d\n", count)
	for i := 1; i <= count; i++ {
		ch <- 10 * i
	}
	close(ch)
}
