package theories

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
