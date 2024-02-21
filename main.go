package main

import "fmt"

func main() {
	type Grades int
	const (
		A Grades = iota // 0
		B               // 1
		C               // 2
		D               // 3
		F               // 4
	)
	fmt.Println(B)
}
