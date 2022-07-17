// Berkay Deniz

// Sieve of Eratosthenes
package main

import (
	"fmt"
)

func sieveOfEratosthenes(number int) {

	// Array of booleans. True indicates the current index is a prime number
	primeNumbers := make([]bool, number+1)

	// Initialize the prime number array as all numbers are true at first.
	for i := 0; i < number+1; i++ {
		primeNumbers[i] = true
	}

	// Mark the first as prime number and mark its factors as non prime numbers
	for i := 2; i*i <= number; i++ {
		if primeNumbers[i] == true {
			for j := i * 2; j <= number; j += i {
				primeNumbers[j] = false
			}
		}

	}

	// The final prime numbers array stores a true value for prime numbers
	// Prints the prime numbers up to the specified number
	fmt.Printf("Prime numbers up to %d : ", number)
	for i := 2; i <= number; i++ {
		if primeNumbers[i] == true {
			fmt.Printf("%d ", i)
		}
	}
	fmt.Println()

}

func main() {
	// Testing the program
	sieveOfEratosthenes(100)
}
