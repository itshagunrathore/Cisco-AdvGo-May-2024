package main

import "fmt"

func main() {
	// invoke "exec" so that it invokes f1
	exec(f1)

	// invoke "exec" so that it invokes f2
	exec(f2)

	exec(func() {
		fmt.Println("anonymous fn invoked")
	})
}

func exec(fn func()) {
	fn()
}

func f1() {
	fmt.Println("f1 invoked")
}

func f2() {
	fmt.Println("f2 invoked")
}
