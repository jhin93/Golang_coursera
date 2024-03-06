# Golang_coursera

findian.go
```go
// Find I, A, N
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Type any word")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	removeSpace := strings.ReplaceAll(input, " ", "")
	fmt.Println("removeSpace : ", removeSpace)

	lowerInput := strings.ToLower(removeSpace)
	fmt.Println("lowerInput : ", lowerInput)

	splitedInput := strings.Split(lowerInput, "")
	fmt.Println("splitedInput : ", splitedInput)

	firstLetter := splitedInput[0]
	lastLetter := splitedInput[len(splitedInput)-1]
	checkLetterA := strings.Contains(lowerInput, "a")

	if firstLetter == "i" && lastLetter == "n" && checkLetterA {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
```

slice.go
```go
// Accomodate Integer, or "X"

func main() {
	// Creating Slice
	intSlice := make([]int, 0, 3)
	// For loop
	for cap(intSlice) > 0 {
		fmt.Println("Type any integer.")
		var inputString string
		// Error checking
		if _, err := fmt.Scan(&inputString); err != nil {
			fmt.Println("Error occurred:", err)
			break
		}
		// If Input is "X", finish the loop
		if inputString == "X" {
			fmt.Println("Exiting the loop.")
			break
		}
		// Checking letter between integer
		if hasLetter(inputString) {
			fmt.Println("Invalid input. Please enter a valid integer.")
			continue
		}
		inputInt, _ := strconv.Atoi(inputString)
		intSlice = append(intSlice, inputInt)
		sort.Ints(intSlice)
		fmt.Println("intSlice:", intSlice)
	}
}

// Checking letter function
func hasLetter(s string) bool {
	for _, char := range s {
		if unicode.IsLetter(char) {
			return true
		}
	}
	return false
}
```

makejson.go
```go
package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// Create map, name, address
	nameAddr := make(map[string]string)
	var inputName string
	var inputAddr string

	// Scan input from terminal
	fmt.Println("Type your name.")
	fmt.Scan(&inputName)
	fmt.Println("Type your address.")
	fmt.Scan(&inputAddr)

	// Put input value into map
	nameAddr["name"] = inputName
	nameAddr["address"] = inputAddr

	jsonData, err := json.Marshal(nameAddr)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Use Marshal() to convert map to json
	fmt.Println(string(jsonData))
}

```

read.go
```go
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
```







