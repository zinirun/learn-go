package theories

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(nums)
	fmt.Println(nums[:3])
	fmt.Println(nums[3:])
	fmt.Println(nums[1:4])
}

func sliceCopy() {
	src := []int{30, 20, 50, 100}
	dest := make([]int, len(src))
	// 1. 직접 복사하는 방법
	for _, i := range src {
		dest[i] = src[i]
	}
	// 2. copy 함수가 있음 -> memmove처럼 겹치는 영역을 복사할때도 쓸 수 있음(효율적)
	copy(dest, src)
	// 3. append 함수 이용
	dest = append([]int(nil), src...)
	// X. dest := src 는 배열 포인터를 복사하므로 deep copy X
}
