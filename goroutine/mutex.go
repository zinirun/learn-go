package goroutine

import (
	"fmt"
	"sync"
)

func MutexExample() {
	myMap := map[string]string{
		"name": "zini",
	}
	// Go언어 기본 패키지 소스코드를 보면 많이 쓰임
	var myRWMutex = &sync.RWMutex{}

	// 쓸 때
	myRWMutex.Lock()
	myMap["job"] = "developer"
	myRWMutex.Unlock()

	// 읽을 때
	myRWMutex.RLock()
	fmt.Println(myMap["name"])
	myRWMutex.RUnlock()
}
