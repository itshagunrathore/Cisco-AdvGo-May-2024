package main

import (
	"fmt"
	"time"
)

func main() {
	// panic("dummy panic")

	go f1() // schedule the execution of f1 through the scheduler (the execution will happen in future)

	f2()

	// block the execution of this function so that the scheduler can look for other goroutines scheduled and execute them

	// POOR man's synchronization techniques. DO NOT USE THESE IN PRODUCTION!!!
	// time.Sleep(300 * time.Millisecond)
	fmt.Scanln()
}

func f1() {
	fmt.Println("f1 started")
	time.Sleep(1 * time.Second)
	fmt.Println("f1 completed")
}

func f2() {
	fmt.Println("f2 invoked")
}
