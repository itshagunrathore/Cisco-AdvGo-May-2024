package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(10)
	go fn(wg)
	wg.Wait()
	fmt.Println("Done")
}

func fn(wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(4 * time.Second)
	fmt.Println("fn invoked")
}
