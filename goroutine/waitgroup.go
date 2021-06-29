package goroutine

import (
	"fmt"
	"sync"
)

// time을 쓰지 않고 WaitGroup 구조체를 사용하여 실행중인 고루틴을 기다리는 방법
var wg = &sync.WaitGroup{}

func runLoopSend2(n int, ch chan int) {
	defer wg.Done()
	for i := 0; i < n; i++ {
		ch <- i
	}
	// 채널이 닫히면 더이상 데이터를 보낼 수 없고 panic이 발생
	close(ch)
}

func runLoopRecv2(ch chan int) {
	defer wg.Done()
	// 채널에 대한 for ... range 문은 채널이 닫힐 때까지 데이터를 받음
	for c := range ch {
		fmt.Println("Received Value: ", c)
	}
}

func ChannelExample2() {
	myChannel := make(chan int)

	wg.Add(2)
	go runLoopSend2(10, myChannel)
	go runLoopRecv2(myChannel)
	wg.Wait()
}
