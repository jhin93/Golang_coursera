package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func GenDisplaceFn(a, v0, s0 float64) func(float64) float64 {
	return func(t float64) float64 {
		result := 0.5*a*math.Pow(t, 2) + v0*t + s0
		return result
	}
}

func main() {
	var float64Slice []float64
	var a, v0, s0, t float64
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Type an acceleration, initial velocity, initial displacement in order. Each value are distinguished by blank")
	scanner.Scan()
	inputLine := scanner.Text()
	if err := scanner.Err(); err != nil {
		fmt.Println("Error occured while scanning typed input", err)
	}
	stringSlice := strings.Fields(inputLine)

	for _, Inputs := range stringSlice {
		eachInt, err := strconv.ParseFloat(Inputs, 64)
		if err != nil {
			fmt.Println("Error occured while convert string to int", err)
		}
		float64Slice = append(float64Slice, eachInt)

	}
	a = float64Slice[0]
	v0 = float64Slice[1]
	s0 = float64Slice[2]

	displaceFn := GenDisplaceFn(a, v0, s0)

	fmt.Println("Enter the time:")
	scanner.Scan()
	timeInput := scanner.Text()
	timeFloat64, err := strconv.ParseFloat(timeInput, 64)
	if err != nil {
		fmt.Println("Error occured while parsing time input to float64", err)
	}
	t = timeFloat64

	destination := displaceFn(t)

	fmt.Println("Final Destination : ", destination)
}
