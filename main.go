package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan bool)

	// 다른 고루틴에서 일정 시간 후에 done 채널에 신호를 보내는 예제
	go func() {
		time.Sleep(2 * time.Second)
		done <- true
	}()

	fmt.Println("Waiting for the signal...")
	// 채널에서 신호를 기다리며, 해당 신호를 변수에 저장하지 않고 그냥 버림
	<-done
	fmt.Println("Received the signal!")
}
