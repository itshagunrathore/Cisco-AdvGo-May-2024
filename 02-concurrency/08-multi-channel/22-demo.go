package main

import (
	"fmt"
	"time"
)

func main() {
	ch := genPrimes()
	for primeNo := range ch {
		fmt.Println("prime no :", primeNo)
	}
	fmt.Println("Done")
}

func genPrimes() <-chan int {
	doneCh := timeOut(10 * time.Second)
	ch := make(chan int)
	go func() {
	OUTER_LOOP:
		for no := 2; ; no++ {
			if isPrime(no) {
				select {
				case <-doneCh:
					break OUTER_LOOP
				case ch <- no:
					time.Sleep(500 * time.Millisecond)
				}
			}
		}
		fmt.Println("all prime number generated")
		close(ch)
	}()
	return ch
}

func timeOut(d time.Duration) <-chan time.Time {
	ch := make(chan time.Time)
	go func() {
		time.Sleep(d)
		close(ch)
	}()
	return ch
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
