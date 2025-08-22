package main

import (
	arithmetic "Task_19Aug/arithmetic"
	"fmt"
)

func main() {
	var x, y int

	fmt.Print("Enter first number: ")
	fmt.Scanln(&x)

	fmt.Print("Enter second number: ")
	fmt.Scanln(&y)

	sum := arithmetic.Add(x, y)
	product := arithmetic.Multiply(x, y)

	// Print results
	fmt.Printf("Sum: %d\n", sum)
	fmt.Printf("Product: %d\n", product)
}
