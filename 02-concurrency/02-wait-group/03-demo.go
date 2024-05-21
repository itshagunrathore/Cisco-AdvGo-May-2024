package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}
	fmt.Println("main started")
	wg.Add(1) // increment the counter by 1
	go fn(wg)
	wg.Wait() // block/wait until the wg counter becomes 0 (default = 0)
	fmt.Println("main completed")
}

func fn(wg *sync.WaitGroup) {
	defer wg.Done() // decrement the counter by 1
	fmt.Println("fn started")
	time.Sleep(3 * time.Second)
	fmt.Println("fn completed")

}
