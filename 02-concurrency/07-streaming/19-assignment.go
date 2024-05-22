package main

func main() {
	// Use the same channel to receive data from genPrimes() & genEvenNos() and print them as and ewhen they are generated
}

func genPrimes(start, end int) {
	// should send one prime number at a time through a channel at 500 ms intervals
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

func genEvenNos(count int) {
	// should send one even number (starting from 0) at a time through a channel at 300 ms intervals
}
