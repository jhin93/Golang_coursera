package main

import (
	"fmt"
	"sync"
)

var i int = 0
var wg sync.WaitGroup
var mut sync.Mutex // 뮤텍스 변수 선언

func inc() {
	mut.Lock() // 뮤텍스로 변수 i의 접근을 잠금
	i = i + 1
	mut.Unlock() // 잠금 해제
	wg.Done()
}

func main() {
	wg.Add(2)
	go inc()
	go inc()
	wg.Wait()
	fmt.Println(i) // 예상 결과는 항상 2가 출력
}
