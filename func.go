// 1개만 리턴할 때
func multiply(a, b int) int { // 인자 => a int, b int
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

// Naked Return & defer
func lenAndUpperNaked(name string) (length int, uppercase string) {
	defer fmt.Println("I'm done") // defer -> function이 끝나고나서 실행
	length, uppercase = len(name), strings.ToUpper(name)
	return // 반환형에서 이미 리턴을 명시해줬음 (변수에 넣기만하면 됨)
}