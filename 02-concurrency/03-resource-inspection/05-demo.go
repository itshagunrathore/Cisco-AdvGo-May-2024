package main

import (
	"flag"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	var count int

	// parsing the command line argument "count"
	flag.IntVar(&count, "count", 0, "# of goroutines to start")
	flag.Parse()

	wg := &sync.WaitGroup{}
	fmt.Println("main started")
	fmt.Printf("Starting %d goroutines.. Hit ENTER to start\n", count)
	fmt.Scanln()
	for i := 1; i <= count; i++ {
		wg.Add(1) // increment the counter by 1
		go fn(wg, i)
	}
	wg.Wait() // block/wait until the wg counter becomes 0 (default = 0)
	fmt.Println("main completed")
}

func fn(wg *sync.WaitGroup, id int) {
	defer wg.Done() // decrement the counter by 1
	fmt.Printf("fn[%d] started\n", id)
	time.Sleep(time.Duration(rand.Intn(20)) * time.Second)
	fmt.Printf("fn[%d] completed\n", id)
}
