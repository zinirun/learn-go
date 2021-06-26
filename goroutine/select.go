package goroutine

import (
	"fmt"
	"time"
)

// select 구문으로 여러 채널을 동시에 제어할 수 있음
// 여러 채널에 데이터를 보내거나 받을 수 있고, 먼저 활성화된 채널의 코드를 실행할 수 있음
// 타임아웃 채널, 수신, 송신 채널을 포함할 수 있음
func SelectExample() {
	ch1, ch2 := make(chan int), make(chan int)
	select {
	case i := <-ch1:
		fmt.Println("Received value on ch1: ", i)
	case ch2 <- 10:
		fmt.Println("Sent value 10 to ch2")
	case <-time.After(1 * time.Second):
		fmt.Println("timed out")
	default:
		fmt.Println("No channel is ready")
	}
}
