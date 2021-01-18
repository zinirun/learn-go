// struct (구조체) -> Typescript랑 비슷
type person struct {
	name    string
	age     int
	favFood []string
}

/* struct 사용
favFood := []string{"ramen", "cheese"}
zini := person{ // value만 적어도 되지만 key를 적어야 명시적임
	name:    "zini",
	age:     25,
	favFood: favFood,
}
fmt.Println(zini)
fmt.Println(zini.name)
*/