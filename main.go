package main

import "fmt"

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct {
	Name       string
	Food       string
	Locomotion string
	Sound      string
}

type Bird struct {
	Name       string
	Food       string
	Locomotion string
	Sound      string
}

type Snake struct {
	Name       string
	Food       string
	Locomotion string
	Sound      string
}

func (c Cow) Eat() {
	fmt.Println(c.Food)
}

func (c Cow) Move() {
	fmt.Println(c.Locomotion)
}

func (c Cow) Speak() {
	fmt.Println(c.Sound)
}

func (b Bird) Eat() {
	fmt.Println(b.Food)
}

func (b Bird) Move() {
	fmt.Println(b.Locomotion)
}

func (b Bird) Speak() {
	fmt.Println(b.Sound)
}

func main() {

}
