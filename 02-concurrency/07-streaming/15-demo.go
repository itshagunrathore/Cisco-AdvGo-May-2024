package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go genNos(ch)
	fmt.Println(<-ch) // (B)(UB)
	fmt.Println(<-ch) // (NB)
	fmt.Println(<-ch) // (B)(UB)
	fmt.Println(<-ch) // (NB)
	fmt.Println(<-ch)
}

func genNos(ch chan<- int) {
	ch <- 10 // (NB)
	time.Sleep(2 * time.Second)
	ch <- 20 // (B)(UB)
	time.Sleep(2 * time.Second)
	ch <- 30 // (NB)
	time.Sleep(2 * time.Second)
	ch <- 40 // (B)
	time.Sleep(2 * time.Second)
	ch <- 50
	time.Sleep(2 * time.Second)
}
