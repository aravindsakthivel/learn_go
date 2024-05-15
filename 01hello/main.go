package main

import "fmt"

// type Shape interface {
// 	area() float64
// 	perimeter() float64
// }

// type rec struct {
// 	length, width float64
// }

// func (r rec) area() float64 {
// 	return r.width * r.length
// }

// func (r rec) perimeter() float64 {
// 	return (2 * r.width) + (2 * r.length)
// }

// func main() {
// 	fmt.Println("Hello, World!")
// 	structLesson()
// 	anonymousStruct()
// 	var vl int = r.area()
// 	fmt.Println(vl)
// 	rec := rec{length: 5, width: 10}
// 	rec.area()
// 	rec.perimeter()
// }

func main() {
	fmt.Println("Hello, World!")
	mainIn()

	var x interface{}

	// Assign an integer value to the interface variable
	x = 42

	// Type assertion to extract the integer value
	val, ok := x.(int)
	if ok {
		fmt.Println("The value of x is an integer:", val)
	} else {
		fmt.Println("x is not an integer.")
	}

}
