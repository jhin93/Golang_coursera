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
	var filename string
	// var names Name
	// var nameSlice []Name
	fmt.Println("Enter the name of the text file : ")
	fmt.Scan(&filename)
	// Read the file
	contents, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	fileContents := string(contents)
	wholeName := strings.Split(fileContents, "\n\n")
	fmt.Println(wholeName[1])
	for _, name := range wholeName {
		fmt.Println(name)
	}
}
