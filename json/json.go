package json

import (
	"encoding/json"
	"fmt"
)

// Marshal (stringify)
type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

func MarshalSample() {
	book := Book{
		Title:  "Learning Go",
		Author: "zini",
	}
	// it converts struct to JSON BYTEARRAY
	byteArray, err := json.Marshal(book)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(byteArray)
	fmt.Println(string(byteArray))

	// Indentation
	byteArray, err = json.MarshalIndent(book, "", "  ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(byteArray))
}

// Unmarshal (parse)
type SensorReading struct {
	Name     string `json:"name"`
	Capacity int    `json:"capacity"`
	Time     string `json:"time"`
}

func UnmarshalSample() {
	jsonString := `{"name": "battery sensor", "capacity": 40, "time": "2019-01-21T19:07:28Z"}`
	var reading SensorReading
	// if it's unstructed data, use map[string]interface{} (Not struct)
	if err := json.Unmarshal([]byte(jsonString), &reading); err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", reading)
}
