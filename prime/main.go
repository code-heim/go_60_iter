package main

import (
	"fmt"
	"iter"
	"math"
)

// Function that returns the next prime number after a given integer
func nextPrime(n int) int {
	if n < 2 {
		return 2
	}
	for {
		n++
		if isPrime(n) {
			return n
		}
	}
}

// Helper function to check if a number is prime
func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// Prime generator using iter.Seq for lazy evaluation
func primeGenerator() iter.Seq[int] {
	num := 1
	seq := func(yield func(int) bool) {
		for {
			num = nextPrime(num)
			if !yield(num) {
				break
			}
		}
	}
	return seq
}

func main() {
	i := 0
	for prime := range primeGenerator() {
		i++
		fmt.Println("Prime #", i, ": ", prime)
		if i == 10 {
			break
		}
	}
}
