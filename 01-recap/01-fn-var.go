package main

import "fmt"

func main() {
	/* Assign a function to a variable */
	/*
		var fn = func() {
			fmt.Println("fn invoked")
		}
		fn()

		var add = func(x, y int) {
			fmt.Println("result :", x+y)
		}
		add(100, 200)

		var subtract = func(x, y int) int {
			return x - y
		}
		var result = subtract(100, 200)
		fmt.Println("subtract result :", result)
	*/

	var fn func()

	fn = func() {
		fmt.Println("fn1 invoked")
	}
	fn()

	fn = func() {
		fmt.Println("fn2 invoked")
	}
	fn()

	var add func(int, int)
	add = func(x, y int) {
		fmt.Println("result :", x+y)
	}
	add(100, 200)

	var subtract func(int, int) int
	subtract = func(x, y int) int {
		return x - y
	}
	var result = subtract(100, 200)
	fmt.Println("subtract result :", result)
}
