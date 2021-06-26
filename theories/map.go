package theories

import "fmt"

//Map: key, value, map[key]value -> value 변수타입을 다르게 하고싶으면 struct 써야함

func UseMap() {
	zini := map[string]string{"name": "zini", "age": "25"}
	for key, value := range zini { // for loop에서 사용할 수도 있음
		fmt.Println(key, value)
	}
	fmt.Println(zini)
}

func UseMap2() {
	//var myMap map[int]string
	myMap := make(map[int]string)
	fmt.Println(myMap)

	// 값이 있는지 확인할 수 있음
	if x, ok := myMap[5]; ok {
		fmt.Println(x)
	}

	// 삭제: delete(mapVariable, key)
	delete(myMap, 4)
}
