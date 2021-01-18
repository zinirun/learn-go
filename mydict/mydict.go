package mydict

import "errors"

// Dictionary type
// struct와 달리 hashmap은 * 가 포함되어 있기때문에 메소드 리시버에서 포인터로 안받아도 됨
type Dictionary map[string]string

var (
	errNotFound        = errors.New("Not Found")
	errWordExists      = errors.New("Taht words already exists")
	errCantUpdate      = errors.New("Can't update for not exists")
	errNothingHappened = errors.New("Nothing happened, non-existing")
)

// Search for a word
/*
	def, err := dictionary.Search("second")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(def)
	}
*/
func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word] // key로 접근하면 value, 존재여부 반환
	if exists {
		return value, nil
	}
	return "", errNotFound
}

// Add a word to the dictionary
/*
	word := "hello"
	def := "Greeting"
	err := dictionary.Add(word, def)
	if err != nil {
		fmt.Println(err)
	}
	hello, _ := dictionary.Search(word)
*/
func (d Dictionary) Add(word, def string) error {
	_, err := d.Search(word)
	switch err {
	case errNotFound:
		d[word] = def
	case nil:
		return errWordExists
	}
	return nil
}

// Update definition of a word
/*
	baseWord := "hello"
	dictionary.Add(baseWord, "first")
	err := dictionary.Update(baseWord, "second")
	if err != nil {
		fmt.Println(err)
	}
	word, _ := dictionary.Search(baseWord)
	fmt.Println(word)
*/
func (d Dictionary) Update(word, newDef string) error {
	_, err := d.Search(word)
	switch err {
	case errNotFound:
		return errCantUpdate
	case nil:
		d[word] = newDef
	}
	return nil
}

// Delete a word
func (d Dictionary) Delete(word string) {
	delete(d, word) // 삭제는 go 자체 지원, key 없으면 nothing happen
}

// CustomDelete - 좀 더 상세한 delete
func (d Dictionary) CustomDelete(word string) error {
	_, err := d.Search(word)
	switch err {
	case errNotFound:
		return errNothingHappened
	case nil:
		delete(d, word)
	}
	return nil
}
