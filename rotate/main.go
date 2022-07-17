// Berkay Deniz

// Rotating a slice

package main

import (
	"fmt"
)

func main() {
	// Example Slice
	numbers := [8]int{1, 2, 3, 4, 5, 6, 7, 8}

	var rotationAmount int
	var rotationSide int // 0 for left, and 1 for right

	fmt.Println("Enter the rotation amount: ")
	fmt.Scanln(&rotationAmount)
	fmt.Println("Enter the rotation side (-1 for left, and 1 for right")
	fmt.Scanln(&rotationSide)

	tempNumbers := [8]int{0, 0, 0, 0, 0, 0, 0, 0}

	for i := 0; i < 8; i++ {
		tempNumbers[(i+rotationSide+8)%8] = numbers[i]
	}

	fmt.Println("Rotated slice is: ")
	fmt.Println(tempNumbers)
}
