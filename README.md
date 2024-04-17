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

2 Go routines, with race condition
```go
package main

import (
	"fmt"
	"sync"
)

var (
	counter int // Shared variable accessed by multiple goroutines.
	wg      sync.WaitGroup
)

// increment function, which modifies the shared counter variable.
// This function is intended to be run in a goroutine.
func increment() {
	defer wg.Done() // Decrement the wait group counter on function exit.
	for i := 0; i < 1000; i++ {
		counter++ // Increment the shared counter.
		// Each goroutine attempts to read, modify, and write back the counter value.
		// Since multiple goroutines modify the counter concurrently without synchronization,
		// it leads to a race condition. A race condition occurs when multiple threads or goroutines
		// access a shared variable concurrently, and at least one of the accesses is a write.
		// This can result in an unpredictable outcome because the order of operations is non-deterministic,
		// leading to the final value of counter being dependent on the order in which the goroutines execute.
	}
}

func main() {
	wg.Add(2) // Set the number of goroutines to wait for.

	go increment() // Launch the first goroutine to increment the counter.
	go increment() // Launch the second goroutine to increment the counter.

	wg.Wait()                              // Wait for all goroutines to finish.
	fmt.Println("Final counter:", counter) // Print the final value of the counter.
	// Due to the race condition, this value might not be what is expected (2000 in this case),
	// because the increments may overwrite each other or not be seen by each goroutine.
}

```

go routine with sorting
```go
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
)

func splitSliceEqually(s []int, parts int) [][]int {
	var result [][]int
	totalLen := len(s)
	baseSize := totalLen / parts
	remainder := totalLen % parts

	start := 0
	for i := 0; i < parts; i++ {
		end := start + baseSize
		if i < remainder {
			end++
		}
		result = append(result, s[start:end])
		start = end
	}
	return result
}

func main() {
	var intSlice []int
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Type at least 4 integers. Each integers are distinguished by blank.")
	scanner.Scan()
	inputLine := scanner.Text()
	// Checking Err
	if err := scanner.Err(); err != nil {
		fmt.Println("Error occur during taking input: ", err)
		return
	}
	stringSlice := strings.Fields(inputLine)

	// Making Slice of integers
	for _, i := range stringSlice {
		num, err := strconv.Atoi(i)
		if err != nil {
			fmt.Println("wrong int during strconv:", i)
			continue
		}
		intSlice = append(intSlice, num)
	}

	splitSlices := splitSliceEqually(intSlice, 4)

	var wg sync.WaitGroup
	for _, split := range splitSlices {
		wg.Add(1)
		go func(s []int) {
			defer wg.Done()
			sort.Ints(s)
		}(split) // go 익명 함수를 split 슬라이스를 인자로 하여 즉시 호출
		fmt.Println("Splited Slice : ", split)
	}
	wg.Wait()

	var mergedSlice []int
	for _, split := range splitSlices {
		mergedSlice = append(mergedSlice, split...)
		sort.Ints(mergedSlice)
	}
	fmt.Println("Result : ", mergedSlice)
}

```

5philo eat concurrently

Implement the dining philosopher’s problem with the following constraints/modifications.

1. There should be 5 philosophers sharing chopsticks, with one chopstick between each adjacent pair of philosophers.

2. Each philosopher should eat only 3 times (not in an infinite loop as we did in lecture)

3. The philosophers pick up the chopsticks in any order, not lowest-numbered first (which we did in lecture).

4. In order to eat, a philosopher must get permission from a host which executes in its own goroutine.

5. The host allows no more than 2 philosophers to eat concurrently.

6. Each philosopher is numbered, 1 through 5.

7. When a philosopher starts eating (after it has obtained necessary locks) it prints “starting to eat <number>” on a line by itself, where <number> is the number of the philosopher.

8. When a philosopher finishes eating (before it has released its locks) it prints “finishing eating <number>” on a line by itself, where <number> is the number of the philosopher.

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

type Chopstick chan bool // 젓가락을 나타내는 채널을 bool 타입으로 정의합니다.

func NewChopstick() Chopstick {
	c := make(Chopstick, 1) // 젓가락 채널을 생성하고,
	c <- true              // 젓가락을 사용 가능하도록 초기화합니다.
	return c                // 생성된 젓가락을 반환합니다.
}

type Philosopher struct { // 철학자를 나타내는 구조체입니다.
	id                  int          // 철학자의 ID
	leftCS, rightCS     Chopstick    // 철학자가 사용하는 두 개의 젓가락
	eatCount            int          // 철학자가 현재까지 식사한 횟수
	maxEatCount         int          // 철학자가 식사할 최대 횟수
	permissionChan, done chan struct{} // 식사 허가를 받고, 식사 완료를 알리는 채널
}

func (p Philosopher) eat(wg *sync.WaitGroup) {
	defer wg.Done() // 식사가 끝날 때 WaitGroup의 카운터를 감소시킵니다.
	for p.eatCount < p.maxEatCount { // 최대 식사 횟수에 도달할 때까지 반복합니다.
		<-p.permissionChan // 호스트로부터 식사를 시작할 수 있는 허가를 받습니다.

		<-p.leftCS // 왼쪽 젓가락을 잡습니다.
		<-p.rightCS // 오른쪽 젓가락을 잡습니다.

		fmt.Printf("starting to eat %d\n", p.id) // 식사 시작을 알립니다.
		time.Sleep(time.Second) // 식사 시간을 1초로 가정합니다.
		fmt.Printf("finishing eating %d\n", p.id) // 식사 완료를 알립니다.

		p.eatCount++ // 식사 횟수를 증가시킵니다.

		p.leftCS <- true // 왼쪽 젓가락을 반환합니다.
		p.rightCS <- true // 오른쪽 젓가락을 반환합니다.

		p.done <- struct{}{} // 식사 완료 신호를 호스트에게 보냅니다.
	}
}

func host(philosophers []*Philosopher) {
	concurrentEaters := make(chan struct{}, 2) // 동시에 식사할 수 있는 철학자 수를 2명으로 제한합니다.

	for _, p := range philosophers { // 모든 철학자에 대해 반복합니다.
		go func(p *Philosopher) { // 각 철학자를 위한 고루틴을 시작합니다.
			for {
				concurrentEaters <- struct{}{} // 동시 식사 슬롯을 점유합니다.
				p.permissionChan <- struct{}{} // 철학자에게 식사 시작 허가를 보냅니다.
				<-p.done // 철학자의 식사 완료 신호를 기다립니다.
				<-concurrentEaters // 동시 식사 슬롯을 반환합니다.
			}
		}(p)
	}
}

func main() {
	wg := &sync.WaitGroup{} // 동기화를 위한 WaitGroup 객체를 생성합니다.
	chopsticks := make([]Chopstick, 5) // 5개의 젓가락을 생성합니다.
	for i := 0; i < 5; i++ {
		chopsticks[i] = NewChopstick() // 각 젓가락을 초기화합니다.
	}

	philosophers := make([]*Philosopher, 5) // 5명의 철학자를 생성합니다.
	for i := 0; i < 5; i++ {
		philosophers[i] = &Philosopher{ // 각 철학자를 초기화합니다.
			id:             i + 1,
			leftCS:         chopsticks[i],
			rightCS:        chopsticks[(i+1)%5],
			maxEatCount:    3,
			permissionChan: make(chan struct{}, 1),
			done:           make(chan struct{}, 1),
		}
		wg.Add(1) // WaitGroup의 카운트를 1 증가시킵니다.
	}

	go host(philosophers) // 호스트 고루틴을 시작합니다.
	for _, p := range philosophers {
		go p.eat(wg) // 각 철학자의 식사 고루틴을 시작합니다.
	}

	wg.Wait() // 모든 철학자가 식사를 마칠 때까지 대기합니다.
}

```