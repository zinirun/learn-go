package theory

// If Else, Variable Expression
func canIDrink(age int) bool {
	if koreanAge := age + 2; koreanAge < 18 { // ;로 변수 선언 후 검사할 수 있음 (조건에만 사용하기위해 변수 선언할 때)
		return false
	}
	return true
}
