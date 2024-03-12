package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var intSlice []int
	var input string
	fmt.Println("Enter no more than 10 integers. Each integers are distinguished by blank.")
	_, err := fmt.Scanln(&input)
	if err != nil {
		fmt.Println("Error occur during taking integer : ", err)
		return
	}
	stringInputs := strings.Fields(input)
	for _, i := range stringInputs {
		num, err := strconv.Atoi(i)
		if err != nil {
			fmt.Println("Error occur during convert Ascii to int : ", err)
		}
		intSlice = append(intSlice, num)
	}
	fmt.Println(intSlice)
}
