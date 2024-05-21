package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	// ver 1.0
	/*
		add(100, 200)
		subtract(100, 200)
	*/

	// ver 2.0
	/*
		logOperation(add, 100, 200)
		logOperation(subtract, 100, 200)
		logOperation(func(x, y int) {
			fmt.Println("Multiply Result : ", x*y)
		}, 100, 200)
	*/

	// ver 3.0
	/*
		var logAdd = getLogOperation(add)
		logAdd(100, 200)

		var logSubtract = getLogOperation(subtract)
		logSubtract(100, 200)
	*/

	/* ver 4.0 */
	/*
		var add = getLogOperation(add)
		var subtract = getLogOperation(subtract)
		// from ver 1.0
		add(100, 200)
		subtract(100, 200)
	*/

	/*
		var add = getProfiledOperation(add)
		add(100, 200)
		var subtract = getProfiledOperation(subtract)
		subtract(100, 200)
	*/

	var logAdd = getLogOperation(add)
	var add = getProfiledOperation(logAdd)
	add(100, 200)

	var logSubtract = getLogOperation(subtract)
	var subtract = getProfiledOperation(logSubtract)
	subtract(100, 200)

}

// ver 1.0
func add(x, y int) {
	fmt.Println("Add Result : ", x+y)
}

func subtract(x, y int) {
	fmt.Println("Subtract Result : ", x-y)
}

// ver 2.0
/*
func logAdd(x, y int) {
	log.Println("Operation started")
	add(x, y)
	log.Println("Operation completed")
}

func logSubtract(x, y int) {
	log.Println("Operation started")
	subtract(x, y)
	log.Println("Operation completed")
}

func logOperation(op func(int, int), x, y int) {
	log.Println("Operation started")
	op(x, y)
	log.Println("Operation completed")
}
*/

// ver 3.0
/*
func getLogOperation(op func(int, int)) func(int, int) {
	return func(x, y int) {
		log.Println("Operation started")
		op(x, y)
		log.Println("Operation completed")
	}
}
*/

// Function Type Declaration
type OperationFn func(int, int)

func getLogOperation(op OperationFn) OperationFn {
	return func(x, y int) {
		log.Println("Operation started")
		op(x, y)
		log.Println("Operation completed")
	}
}

func getProfiledOperation(op OperationFn) OperationFn {
	return func(x, y int) {
		start := time.Now()
		op(x, y)
		elapsed := time.Since(start)
		fmt.Println("Elapsed :", elapsed)
	}
}
