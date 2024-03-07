package main

import "fmt"

func foo(sli []int) []int { // slice contain a pointer to array
	sli[0] = sli[0] + 1
	return sli
}

func main() {
	a := []int{1, 2, 3}
	fmt.Print(a) // [1 2 3]
	foo(a)
	fmt.Print(a) // [2 2 3]
}
