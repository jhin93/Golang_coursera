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

Bubble Sort 
```go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var intSlice []int
	// Making scanner object
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Type no more than 10 integers. Each integers are distinguished by blank.")
	// Scan the input that the user typed
	scanner.Scan()
	// Read the contents of input
	inputLine := scanner.Text()
	// Checking Err
	if err := scanner.Err(); err != nil {
		fmt.Println("Error occur during taking input: ", err)
		return
	}
	// Making Slice of integers
	stringSlice := strings.Fields(inputLine)

	for _, eachInt := range stringSlice {
		num, err := strconv.Atoi(eachInt)
		if err != nil {
			fmt.Println("Error occur during converting string to int: ", err)
			continue
		}
		intSlice = append(intSlice, num)
	}
	// Checking whether the length of the slice over 10
	if len(intSlice) > 10 {
		fmt.Println("Do not type more than 10 integers. Try again")
		return
	}
	BubbleSort(intSlice)
	fmt.Println("Sorted slice of integers : ", intSlice)
}

func BubbleSort(slice []int) {
	for i := 0; i < len(slice); i++ {
		for j := 0; j < len(slice)-i-1; j++ {
			if slice[j] > slice[j+1] {
				Swap(slice, j)
			}
		}
	}
}

func Swap(slice []int, i int) {
	slice[i], slice[i+1] = slice[i+1], slice[i]
}

```

GenDisplaceFn(return func)
```go
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

```

switch animal and actions
```go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func (a Animal) Eat() {
	fmt.Println(a.food)
}

func (a Animal) Move() {
	fmt.Println(a.locomotion)
}

func (a Animal) Speak() {
	fmt.Println(a.noise)
}

func main() {
	cow := Animal{"grass", "walk", "moo"}
	bird := Animal{"worms", "fly", "peep"}
	snake := Animal{"mice", "slither", "hsss"}

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println(">")
		scanner.Scan()
		inputLine := scanner.Text()
		if err := scanner.Err(); err != nil {
			fmt.Println("Error occured while scanning typed input", err)
		}
		stringSlice := strings.Fields(inputLine)

		if len(stringSlice) != 2 {
			fmt.Println("Invalid Input. Please enter two words.")
			return
		}
		switch stringSlice[0] {
		case "cow":
			switch stringSlice[1] {
			case "food":
				cow.Eat()
			case "move":
				cow.Move()
			case "speak":
				cow.Speak()
			default:
				fmt.Println("Type among 3 actions(food, move, speak)")
			}
		case "bird":
			switch stringSlice[1] {
			case "food":
				bird.Eat()
			case "move":
				bird.Move()
			case "speak":
				bird.Speak()
			default:
				fmt.Println("Type among 3 actions(food, move, speak)")
			}
		case "snake":
			switch stringSlice[1] {
			case "food":
				snake.Eat()
			case "move":
				snake.Move()
			case "speak":
				snake.Speak()
			default:
				fmt.Println("Type among 3 actions(food, move, speak)")
			}
		default:
			fmt.Println("Type among 3 animals(cow, bird, snakes)")
		}
	}

}

```

Make newanimal and actions
```go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type DefaultAnimal struct {
	Name       string
	Food       string
	Locomotion string
	Sound      string
}

type Cow struct {
	DefaultAnimal
}

type Bird struct {
	DefaultAnimal
}

type Snake struct {
	DefaultAnimal
}

func (a DefaultAnimal) Eat() {
	fmt.Println(a.Food)
}

func (a DefaultAnimal) Move() {
	fmt.Println(a.Locomotion)
}

func (a DefaultAnimal) Speak() {
	fmt.Println(a.Sound)
}
func main() {
	newAnimal := make(map[string]Animal)
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println(">")
		scanner.Scan()
		inputLine := scanner.Text()
		if err := scanner.Err(); err != nil {
			fmt.Println("Error occured while scanning typed input", err)
		}
		stringSlice := strings.Fields(inputLine)

		if len(stringSlice) != 3 {
			fmt.Println("Invalid Input. Please enter at least three words. First word must be 'newanimal' or 'query'.")
			return
		}

		switch stringSlice[0] {
		case "newanimal":
			switch stringSlice[2] {
			case "cow":
				newAnimal[stringSlice[1]] = Cow{DefaultAnimal{Name: stringSlice[1], Food: "grass", Locomotion: "walk", Sound: "moo"}}
				fmt.Println("Created it!")
			case "bird":
				newAnimal[stringSlice[1]] = Bird{DefaultAnimal{Name: stringSlice[1], Food: "worms", Locomotion: "fly", Sound: "peep"}}
				fmt.Println("Created it!")
			case "snake":
				newAnimal[stringSlice[1]] = Bird{DefaultAnimal{Name: stringSlice[1], Food: "mice", Locomotion: "slither", Sound: "hsss"}}
				fmt.Println("Created it!")
			default:
				fmt.Println("Invalid animal type. Please enter 'cow', 'bird', or 'snake'.")
			}

		case "query":
			animal, ok := newAnimal[stringSlice[1]]
			if !ok {
				fmt.Println("Animal not found. Try again.")
				continue
			}
			switch stringSlice[2] {
			case "eat":
				animal.Eat()
			case "move":
				animal.Move()
			case "speak":
				animal.Speak()
			default:
				fmt.Println("Invalid action. Please enter 'eat', 'move', or 'speak'.")
			}
		default:
			fmt.Println("First word must be 'newanimal' or 'query'.")
		}
	}
}

```


