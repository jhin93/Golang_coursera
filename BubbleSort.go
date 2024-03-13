package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var intSlice []int
	// Making scanner object
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Type no more than 10 integers. Each integers are distinguished by blank.")
	// Scan the input that the user typed
	scanner.Scan()
	// Read the contents of input
	inputLine := scanner.Text()
	// Checking Err
	if err := scanner.Err(); err != nil {
		fmt.Println("Error occur during taking input: ", err)
		return
	}
	// Making Slice of integers
	stringSlice := strings.Fields(inputLine)

	for _, eachInt := range stringSlice {
		num, err := strconv.Atoi(eachInt)
		if err != nil {
			fmt.Println("Error occur during converting string to int: ", err)
			continue
		}
		intSlice = append(intSlice, num)
	}
	// Checking whether the length of the slice over 10
	if len(intSlice) > 10 {
		fmt.Println("Do not type more than 10 integers. Try again")
		return
	}

}
