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

	for i := 0; i < len(wholeName); i++ {
		fullName := strings.Split(wholeName[i], " ")
		names.fname = fullName[0]
		names.lname = fullName[1]
		namesSlice = append(namesSlice, names)
	}
	fmt.Println("namesSlice : ", namesSlice)
}
