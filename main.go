package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done() // 고루틴이 종료될 때마다 WaitGroup의 카운터를 감소시킵니다.
	fmt.Printf("Worker %d 시작\n", id)
	time.Sleep(time.Second) // 임의의 작업을 수행하는 시간 대신 sleep을 사용합니다.
	fmt.Printf("Worker %d 완료\n", id)
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1) // 고루틴을 시작하기 전에 WaitGroup의 카운터를 증가시킵니다.
		go worker(i, &wg)
	}

	wg.Wait() // 모든 고루틴이 종료될 때까지 메인 고루틴을 대기시킵니다.

	fmt.Println("모든 작업 완료")
}
