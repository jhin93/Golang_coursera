package main

import "fmt"

func main() {
	var input float64
	fmt.Println("Type any floating number")
	// Allocate input to variable
	fmt.Scan(&input)
	// Truncate to integer
	result := int(input)
	// Print result
	fmt.Println("Truncated integer :", result)
}
