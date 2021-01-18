// Low Level Programming
// 메모리개념은 c와 같음

func lowlevel() {
	a := 2
	b := a
	a = 10 // b에는 영향이 없다
	fmt.Println(a, b)

	a := 2
	b := &a
	a = 10
	*b = 20 // b로 a의 값을 바꿀수도 있음
	fmt.Println(&a, a, b, *b)
}