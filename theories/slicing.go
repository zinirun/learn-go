package theories

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(nums)
	fmt.Println(nums[:3])
	fmt.Println(nums[3:])
	fmt.Println(nums[1:4])
}
