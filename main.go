package main

import (
	"fmt"
)

func main() {
	// 용량이 2인 버퍼드 채널 생성
	c := make(chan int, 2)

	// 채널에 데이터를 가득 채움
	c <- 1
	c <- 2

	// 채널에 데이터가 가득 차 있으므로 송신자가 블록됨
	// 이후 코드는 실행되지 않음
	c <- 3 // 이부분에서 에러 발생

	// 메인 함수에서는 채널에서 데이터를 받아 출력
	fmt.Println(<-c)
	fmt.Println(<-c)
}
