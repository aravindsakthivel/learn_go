package main

import "fmt"

// Describer interface
type Describer interface {
	Describe() string
}

// Person struct
type Person struct {
	Name string
	Age  int
}

// Describe method implements Describer
func (p Person) Describe() string {
	return fmt.Sprintf("%s is %d years old", p.Name, p.Age)
}

// Function accepting Describer interface
func printDescription(d Describer) {
	fmt.Println(d.Describe())
}

// Main function
func main() {
	// Create a new Person instance
	p := Person{Name: "Alice", Age: 30}

	// Call the printDescription function with a Person
	printDescription(p)
}
