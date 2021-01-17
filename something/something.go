package something

import "fmt"

// 메소드가 대문자로 시작하는 이유: export하려면 첫 문자가 대문자 (자동으로 export)
func SayHello() {
	fmt.Println("Hello")
}

func SayBye() {
	fmt.Println("Bye")
}
