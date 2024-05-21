package main

import (
	"fmt"
	"sync"
)

// concurrent safe counter
type Counter struct {
	sync.Mutex
	count int
}

func (c *Counter) Increment() {
	c.Lock()
	{
		c.count++
	}
	c.Unlock()
}

var c Counter

func main() {
	wg := &sync.WaitGroup{}
	for i := 0; i < 200; i++ {
		wg.Add(1)
		go increment(wg)
	}
	wg.Wait()
	fmt.Printf("count = %d\n", c.count)
}

func increment(wg *sync.WaitGroup) {
	defer wg.Done()
	c.Increment()
}
