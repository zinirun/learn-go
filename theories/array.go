package theories

import "fmt"

// Array: [size]type{...}
// size를 지정하지 않는 이상 초기화한 인덱스의 범위를 벗어나서 직접 변경 불가능 -> append를 사용함 (slice로 배열을 직접 바꾸는게 아닌 것에 유의)
func ArraySample() {
	names := [5]string{"zini", "run", "exciting"}
	names[3] = "hello"
	names[4] = "world"
	// names[5] = "dream" -> ERROR

	unlimited := []string{"zini", "world"}
	// unlimited[2] = "kkk" -> ERROR, []로 배열 선언하면 append 써야함
	unlimited = append(unlimited, "golang") // append는 직접 수정하지 않고 새로운 slice를 리턴함
	fmt.Println(names, unlimited)

	subSlice := make([]string, 5)
	copy(subSlice, names[:])
}
