package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

type ChopS struct {
	sync.Mutex
}

type Philo struct {
	leftCS, rightCS *ChopS
}

func (p Philo) eat(wg *sync.WaitGroup) {
	// 3번만 먹는다는 조건 필요
	// 3번 다 먹으면 defer wg.Done() 실행으로 waitGroup 감소
	defer wg.Done()
	for i := 0; i < 3; i++ {
		p.leftCS.Lock()
		p.rightCS.Lock()

		fmt.Println("eating")

		p.rightCS.Unlock()
		p.leftCS.Unlock()

		fmt.Println("finish")
	}
}

func main() {
	wg.Add(5)

	// Making 5 Chopsticks
	CSticks := make([]*ChopS, 5)
	for i := 0; i < 5; i++ {
		CSticks[i] = new(ChopS)
	}

	// Making 5 Philosophers
	philos := make([]*Philo, 5)
	for i := 0; i < 5; i++ {
		philos[i] = &Philo{leftCS: CSticks[i], rightCS: CSticks[(i+1)%5]}
	}

	for _, philo := range philos {
		go philo.eat(&wg)
	}

	wg.Wait()

}
