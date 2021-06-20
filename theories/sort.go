package theories

import (
	"fmt"
	"sort"
)

func SortSample() {
	myInts := []int{1, 3, 2, 6, 5, 4}
	sort.Ints(myInts)

	fmt.Println(myInts)
}
