package main

import (
	"fmt"
	"sync"
	"time"
)

type Chopstick chan bool

func NewChopstick() Chopstick {
	c := make(Chopstick, 1)
	c <- true // Initialize as available
	return c
}

type Philosopher struct {
	id                   int
	leftCS, rightCS      Chopstick
	eatCount             int
	maxEatCount          int
	permissionChan, done chan struct{}
}

func (p Philosopher) eat(wg *sync.WaitGroup) {
	defer wg.Done()
	for p.eatCount < p.maxEatCount {
		<-p.permissionChan // Wait for permission to eat

		<-p.leftCS  // Take left chopstick
		<-p.rightCS // Take right chopstick

		fmt.Printf("starting to eat %d\n", p.id)
		time.Sleep(time.Second) // Simulate eating
		fmt.Printf("finishing eating %d\n", p.id)

		p.eatCount++

		p.leftCS <- true  // Put down left chopstick
		p.rightCS <- true // Put down right chopstick

		p.done <- struct{}{} // Signal completion
	}
}

func host(philosophers []*Philosopher) {
	concurrentEaters := make(chan struct{}, 2) // Allow 2 philosophers to eat at the same time

	for _, p := range philosophers {
		go func(p *Philosopher) {
			for {
				concurrentEaters <- struct{}{} // Block if already two are eating
				p.permissionChan <- struct{}{}
				<-p.done
				<-concurrentEaters
			}
		}(p)
	}
}

func main() {
	wg := &sync.WaitGroup{}
	chopsticks := make([]Chopstick, 5)
	for i := 0; i < 5; i++ {
		chopsticks[i] = NewChopstick()
	}

	philosophers := make([]*Philosopher, 5)
	for i := 0; i < 5; i++ {
		philosophers[i] = &Philosopher{
			id:             i + 1,
			leftCS:         chopsticks[i],
			rightCS:        chopsticks[(i+1)%5],
			maxEatCount:    3,
			permissionChan: make(chan struct{}, 1),
			done:           make(chan struct{}, 1),
		}
		wg.Add(1)
	}

	go host(philosophers)
	for _, p := range philosophers {
		go p.eat(wg)
	}

	wg.Wait()
}
