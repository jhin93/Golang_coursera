package main

import "fmt"

func main() {
	var x int = 1
	var y int
	var ip *int

	ip = &x
	y = *ip
	fmt.Println(y) // 1
}
