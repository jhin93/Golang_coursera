package main

import "fmt"

func main() {
	// new() function creates a variable and returns a pointer to the variable
	ptr := new(int) // 0xc000012088
	*ptr = 3
	fmt.Println(*ptr) // 3
}
