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
