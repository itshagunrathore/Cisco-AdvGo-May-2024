package main

import (
	"fmt"
	"sync"
	"time"
)

type ResultType struct {
	No   int
	Type string
}

func main() {
	// Use the same channel to receive data from genPrimes() & genEvenNos() and print them as and ewhen they are generated
	wg := &sync.WaitGroup{}
	ch := make(chan ResultType)
	// Ver 1.0
	/*
		go func() {
			wg.Add(1)
			go genPrimes(2, 100, ch, wg)
			wg.Add(1)
			go genEvenNos(60, ch, wg)
			wg.Wait()
			close(ch)
		}()
		for result := range ch {
			fmt.Printf("Type : %q, No : %d\n", result.Type, result.No)
		}
	*/

	// Ver 2.0
	done := make(chan struct{})
	go func() {
		for result := range ch {
			fmt.Printf("Type : %q, No : %d\n", result.Type, result.No)
		}
		fmt.Println("All the result printed")
		close(done)
	}()
	wg.Add(1)
	go genPrimes(2, 50, ch, wg)
	wg.Add(1)
	go genEvenNos(30, ch, wg)
	wg.Wait()
	close(ch)
	<-done
}

func genPrimes(start, end int, ch chan<- ResultType, wg *sync.WaitGroup) {
	defer wg.Done()
	for no := start; no <= end; no++ {
		if isPrime(no) {
			ch <- ResultType{No: no, Type: "Prime"}
			time.Sleep(500 * time.Millisecond)
		}
	}
	fmt.Println("all prime number generated")
}

// utility function to check if a number is prime
func isPrime(no int) bool {
	for i := 2; i <= (no / 2); i++ {
		if no%i == 0 {
			return false
		}
	}
	return true
}

func genEvenNos(count int, ch chan<- ResultType, wg *sync.WaitGroup) {
	defer wg.Done()
	// should send one even number (starting from 0) at a time through a channel at 300 ms intervals
	for i, no := 0, 0; i < count; i, no = i+1, no+2 {
		ch <- ResultType{No: no, Type: "Even"}
		time.Sleep(300 * time.Millisecond)
	}
	fmt.Println("all even number generated")
}
