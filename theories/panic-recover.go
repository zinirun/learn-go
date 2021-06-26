package theories

import "fmt"

func checkPanic() {
	if r := recover(); r != nil {
		fmt.Println("Panic was captured: ", r)
	}
}

func panicTest(p bool) {
	defer checkPanic()
	if p {
		panic("panic requested")
	}
}

func PanicExample() {
	panicTest(true)
	fmt.Println("Hello GO")
}
