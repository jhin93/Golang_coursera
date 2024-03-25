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
