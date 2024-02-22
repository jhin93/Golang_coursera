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

	input := args[1]

	number, err := trunc.Trunc(input)
	if err != nil {
		return // 왜 여기에 return 0 하면 에러가 나는가
	}

	fmt.Println("Truncated number : ", number)
}
