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
	multiplier, divisor := 100, 7
	q, r := divide(multiplier, divisor)
	fmt.Printf("Dividing %d by %d, quotinet = %d and remainder = %d\n", multiplier, divisor, q, r)

}

// 3rd party code
func divide(x, y int) (quotinet, remainder int) {
	if y == 0 {
		panic(errors.New("divisor cannot be zero"))
	}
	quotinet, remainder = x/y, x%y
	return
}
