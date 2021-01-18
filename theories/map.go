package theory

import "fmt"

//Map: key, value, map[key]value -> value 변수타입을 다르게 하고싶으면 struct 써야함

func useMap() {
	zini := map[string]string{"name": "zini", "age": "25"}
	for key, value := range zini { // for loop에서 사용할 수도 있음
		fmt.Println(key, value)
	}
	fmt.Println(zini)
}
