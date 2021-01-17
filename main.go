package main

// Formatting을 위한 라이브러리
import (
	"fmt"
	"strings"
)

func multiply(a, b int) int {
	return a * b
}

// 두가지를 리턴할 수도 있음 (쓰는 쪽에서도 모든 리턴값 복수로 받아야함)
func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

// 인자를 무제한으로 받을 수도 있음 -> repeatMe("zini", "is", "best", "developer") -> [ ... , ... , ... ]
func repeatMe(words ...string) {
	fmt.Println(words)
}

// Naked Return
func lenAndUpperNaked(name string) (length int, uppercase string) {
	defer fmt.Println("I'm done") // defer -> function이 끝나고나서 실행
	length, uppercase = len(name), strings.ToUpper(name)
	return // 반환형에서 이미 리턴을 명시해줬음 (변수에 넣기만하면 됨)
}

// For Loop
func superAdd(numbers ...int) int {
	total := 0
	for _, number := range numbers { // index, value
		total += number
	}
	/*
		for i:=0; i<len(numbers); i++ {
			fmt.Println(i, numbers[i])
		}
	*/
	return total
}

// If Else, Variable Expression
func canIDrink(age int) bool {
	if koreanAge := age + 2; koreanAge < 18 { // ;로 변수 선언 후 검사할 수 있음 (조건에만 사용하기위해 변수 선언할 때)
		return false
	}
	return true
}

// Switch
func canIDrinkSwitch(age int) bool {
	switch { // switch var { .. case 조건: .. } 로 사용해도 됨, 마찬가지로 variable expression; var도 가능
	case age < 18:
		return false
	case age == 18:
		return true
	case age > 50:
		return false
	}
	return false
}

// struct (구조체) -> Typescript랑 비슷
type person struct {
	name    string
	age     int
	favFood []string
}

/* struct 사용
favFood := []string{"ramen", "cheese"}
zini := person{ // value만 적어도 되지만 key를 적어야 명시적임
	name:    "zini",
	age:     25,
	favFood: favFood,
}
fmt.Println(zini)
fmt.Println(zini.name)
*/

// GO의 시작점 func main
/*
var name string = "zini" (상수는 const)
name := "zini" // Go가 별도로 타입선언 안해도 타입을 찾음 (축약형은 func 안에서만 가능)
name = "run"
*/
func main() {

}

/* 배열

names := [5]string{"zini", "run", "exciting"}
names[3] = "hello"
names[4] = "world"
// names[5] = "dream" -> ERROR

unlimited := []string{"zini", "world"}
// unlimited[2] = "kkk" -> ERROR, []로 배열 선언하면 append 써야함
unlimited = append(unlimited, "golang") // append는 직접 수정하지 않고 새로운 slice를 리턴함
fmt.Println(names, unlimited)

*/

/* Map: key, value, map[key]value -> value 변수타입을 다르게 하고싶으면 struct 써야함

zini := map[string]string{"name": "zini", "age": "25"}
for key, value := range zini { // for loop에서 사용할 수도 있음
	fmt.Println(key, value)
}
fmt.Println(zini)

*/

/* Low Level Programming

// 메모리개념은 c와 같음
a := 2
b := a
a = 10 // b에는 영향이 없다
fmt.Println(a, b)

a := 2
b := &a
a = 10
*b = 20 // b로 a의 값을 바꿀수도 있음
fmt.Println(&a, a, b, *b)

*/
