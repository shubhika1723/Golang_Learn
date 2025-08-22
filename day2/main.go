package main

import (
	"day2/util"
	"fmt"
)

func main() {
	fmt.Printf("add(5,3) = %d\n", util.Add(5, 3))
	fmt.Printf("Subtract(5,3) = %d\n", util.Subtract(5, 3))
	fmt.Printf("Multiply(5,3) = %d\n", util.Multiply(5, 3))
	fmt.Printf("Division(5,3) = %.2f\n", util.Division(5, 3))
}
