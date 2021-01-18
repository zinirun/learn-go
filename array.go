func array() {
	names := [5]string{"zini", "run", "exciting"}
	names[3] = "hello"
	names[4] = "world"
	// names[5] = "dream" -> ERROR

	unlimited := []string{"zini", "world"}
	// unlimited[2] = "kkk" -> ERROR, []로 배열 선언하면 append 써야함
	unlimited = append(unlimited, "golang") // append는 직접 수정하지 않고 새로운 slice를 리턴함
	fmt.Println(names, unlimited)
}