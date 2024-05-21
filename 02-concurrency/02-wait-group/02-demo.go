package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("main started")
	wg.Add(1) // increment the counter by 1
	go fn()
	wg.Wait() // block/wait until the wg counter becomes 0 (default = 0)
	fmt.Println("main completed")
}

func fn() {
	fmt.Println("fn started")
	time.Sleep(3 * time.Second)
	fmt.Println("fn completed")
	wg.Done() // decrement the counter by 1
}
