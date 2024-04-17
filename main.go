package main

import (
	"fmt"
	"sync"
	"time"
)

type Chopstick chan bool // 젓가락을 나타내는 채널을 bool 타입으로 정의합니다.

func NewChopstick() Chopstick {
	c := make(Chopstick, 1) // 젓가락 채널을 생성하고,
	c <- true               // 젓가락을 사용 가능하도록 초기화합니다.
	return c                // 생성된 젓가락을 반환합니다.
}

type Philosopher struct { // 철학자를 나타내는 구조체입니다.
	id                   int           // 철학자의 ID
	leftCS, rightCS      Chopstick     // 철학자가 사용하는 두 개의 젓가락
	eatCount             int           // 철학자가 현재까지 식사한 횟수
	maxEatCount          int           // 철학자가 식사할 최대 횟수
	permissionChan, done chan struct{} // 식사 허가를 받고, 식사 완료를 알리는 채널
}

func (p Philosopher) eat(wg *sync.WaitGroup) {
	defer wg.Done()                  // 식사가 끝날 때 WaitGroup의 카운터를 감소시킵니다.
	for p.eatCount < p.maxEatCount { // 최대 식사 횟수에 도달할 때까지 반복합니다.
		<-p.permissionChan // 호스트로부터 식사를 시작할 수 있는 허가를 받습니다.
		<-p.leftCS         // 왼쪽 젓가락을 잡습니다. <-p.leftCS 가 가능한 이유는 NewChopstick() 함수로 초기화가 됐기 때문
		<-p.rightCS        // 오른쪽 젓가락을 잡습니다. <-p.rightCS 가 가능한 이유는 NewChopstick() 함수로 초기화가 됐기 때문

		fmt.Printf("starting to eat %d\n", p.id)  // 식사 시작을 알립니다.
		time.Sleep(time.Second)                   // 식사 시간을 1초로 가정합니다.
		fmt.Printf("finishing eating %d\n", p.id) // 식사 완료를 알립니다.

		p.eatCount++ // 식사 횟수를 증가시킵니다.

		p.leftCS <- true  // 왼쪽 젓가락을 반환합니다.
		p.rightCS <- true // 오른쪽 젓가락을 반환합니다.

		p.done <- struct{}{} // 식사 완료 신호를 호스트에게 보냅니다.
	}
}

func host(philosophers []*Philosopher) {
	concurrentEaters := make(chan struct{}, 2) // 동시에 식사할 수 있는 철학자 수를 2명으로 제한합니다.

	for _, p := range philosophers { // 모든 철학자에 대해 반복합니다.
		go func(p *Philosopher) { // 각 철학자를 위한 고루틴을 시작합니다.
			for {
				concurrentEaters <- struct{}{} // 동시 식사 슬롯을 점유합니다. 'struct{}'는 타입, '{}'는 인스턴스. 'struct{}' 까지만 작성하면 정의일 뿐, 실제 메모리에 할당된 구조체가 아님. 'struct{}{}' 패턴은 데이터를 전송하지 않고 신호만을 전달하겠다는 의도.
				p.permissionChan <- struct{}{} // 철학자에게 식사 시작 허가를 보냅니다.
				// 이 시점에서 eat() 함수가 발동하고, 채널로 식사 마쳤다는 신호를 받음.
				<-p.done           // 철학자의 식사 완료 신호를 기다립니다. eat() 함수가 p.done <- struct{}{} 으로 식사가 끝났다는 신호르 보내는 것을 받는 구문.
				<-concurrentEaters // 동시 식사 슬롯을 반환합니다.
			}
		}(p)
	}
}

func main() {
	wg := &sync.WaitGroup{}            // 동기화를 위한 WaitGroup 객체를 생성합니다.
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
