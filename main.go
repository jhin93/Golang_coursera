package main

import (
	"Golang_coursera/trunc"
	"fmt"
	"os"
)

func main() {
	// Get Input
	args := os.Args
	if len(args) < 2 {
		fmt.Println("Type floating point number")
	}
	// Declare input variable
	input := args[1]

	// Call Trunc func in trunc.go
	number, err := trunc.Trunc(input)
	if err != nil {
		return
	}
	// Print the result
	fmt.Println("Truncated number : ", number)
}
