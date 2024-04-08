package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func splitSliceEqually(s []int, parts int) [][]int {
	var result [][]int
	totalLen := len(s)
	baseSize := totalLen / parts
	remainder := totalLen % parts

	start := 0
	for i := 0; i < parts; i++ {
		end := start + baseSize
		if i < remainder {
			end++
		}
		result = append(result, s[start:end])
		start = end
	}
	return result
}

func main() {
	var intSlice []int
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Type at least 4 integers. Each integers are distinguished by blank.")
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

	splitSlices := splitSliceEqually(intSlice, 4)

	fmt.Println("splitSlices:", splitSlices)
	for _, split := range splitSlices {
		sort.Ints(split)
		fmt.Println(split)
	}
}
