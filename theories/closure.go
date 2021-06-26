package theories

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func ClosureExample() {
	sumClosure := adder() // sum = 0
	sumClosure(1)         // sum = 1
	sumClosure(2)         // sum = 3
}
