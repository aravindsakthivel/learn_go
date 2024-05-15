package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// Define a struct that matches the structure of your JSON data
type Person struct {
	Name string
	Age  int
}

func main() {
	// Example JSON data
	jsonData := `{"name": "John", "age": 30}`

	// Declare a variable to store the unmarshaled data
	var p Person

	// Unmarshal the JSON data into the struct
	err := json.Unmarshal([]byte(jsonData), &p)
	if err != nil {
		log.Fatal(err)
	}

	// Print the struct
	// fmt.Printf("Name: %s, Age: %d\n", p.Name, p.Age)
	fmt.Println(p)
}
