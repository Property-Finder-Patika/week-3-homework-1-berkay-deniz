// Berkay Deniz

// Prime Factorization
package main

import "fmt"

func main() {

	var number int

	// Take the input from the user
	fmt.Print("Enter the number: ")
	fmt.Scanln(&number)

	fmt.Println("Factors of ", number, " are: ")

	// Iterate all numbers up to the input number and check for divisibility
	var counter int
	for counter = 1; counter <= number; counter++ {
		// If a factor is found, print it
		if number%counter == 0 {
			fmt.Println(counter)
		}
	}
}
