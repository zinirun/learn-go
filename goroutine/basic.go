package goroutine

import (
	"fmt"
	"time"
)

func runLoopSend(n int, ch chan int) {
	for i := 0; i < n; i++ {
		ch <- i
	}
	// 채널이 닫히면 더이상 데이터를 보낼 수 없고 panic이 발생
	close(ch)
}

func runLoopRecv(ch chan int) {
	// 채널에 대한 for ... range 문은 채널이 닫힐 때까지 데이터를 받음
	for c := range ch {
		fmt.Println("Received Value: ", c)
	}
}

func ChannelExample() {
	myChannel := make(chan int)
	go runLoopSend(10, myChannel)
	go runLoopRecv(myChannel)
	time.Sleep(2 * time.Second)
}
