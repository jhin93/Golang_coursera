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
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Type series of integers. Each integers are distinguished by blank.")
	scanner.Scan()
	inputLine := scanner.Text()
	// Checking Err
	if err := scanner.Err(); err != nil {
		fmt.Println("Error occur during taking input: ", err)
		return
	}
	stringSlice := strings.Fields(inputLine)

	// Making Slice of integers
	for _, i := range stringSlice {
		num, err := strconv.Atoi(i)
		if err != nil {
			fmt.Println("wrong int during strconv:", i)
			continue
		}
		intSlice = append(intSlice, num)
	}

	fmt.Println("typed integer:", intSlice)
}
