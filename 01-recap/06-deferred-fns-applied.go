package main

import (
	"errors"
	"fmt"
)

func main() {
	defer func() {
		fmt.Println("	[main] deferred")
		if err := recover(); err != nil {
			fmt.Println("	[main] app panicked : ", err)
			return
		}
		fmt.Println("	[main] thank you!")
	}()
	/*
		multiplier := 100
		var divisor int
		fmt.Print("Enter the divisor :")
		fmt.Scanln(&divisor)
		q, r := divide(multiplier, divisor)
		fmt.Printf("Dividing %d by %d, quotinet = %d and remainder = %d\n", multiplier, divisor, q, r)
	*/

	// Using divideWrapper
	for {
		multiplier := 100
		var divisor int
		fmt.Print("Enter the divisor :")
		fmt.Scanln(&divisor)
		q, r, err := divideWrapper(multiplier, divisor)
		if err != nil {
			fmt.Println("Error :", err)
			continue
		}
		fmt.Printf("Dividing %d by %d, quotinet = %d and remainder = %d\n", multiplier, divisor, q, r)
		break
	}

}

// wrapper for 'divide' to convert the 'panic' into an 'error'
func divideWrapper(x, y int) (quotinet, remainder int, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
			return
		}
	}()
	quotinet, remainder = divide(x, y)
	return
}

// 3rd party code
func divide(x, y int) (quotinet, remainder int) {
	if y == 0 {
		panic(errors.New("divisor cannot be zero"))
	}
	quotinet, remainder = x/y, x%y
	return
}
