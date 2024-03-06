package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type Name struct {
	fname string
	lname string
}

func main() {
	// Declare variables
	var filename string
	var names Name
	var namesSlice []Name
	fmt.Println("Enter the name of the text file : ")
	fmt.Scan(&filename)
	// Read the file
	contents, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	fileContents := string(contents)
	wholeName := strings.Split(fileContents, "\n")
	// Make slice of structs
	for i := 0; i < len(wholeName); i++ {
		fullName := strings.Split(wholeName[i], " ")
		names.fname = fullName[0][:min(len(fullName[0]), 20)]
		names.lname = fullName[1][:min(len(fullName[1]), 20)]
		namesSlice = append(namesSlice, names)
	}
	// Print first and last names found in each struct
	for _, value := range namesSlice {
		fmt.Printf("%s, %s\n", value.fname, value.lname)
	}
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
