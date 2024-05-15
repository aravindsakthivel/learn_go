// package main

// // create an interface with findArea method

// type IShape interface {
// 	getArea() float64
// }

// type SRectangle struct {
// 	length, width float64
// }

// func (r SRectangle) getArea() float64 {
// 	return r.width * r.length
// }

// // to check the assert and getArea let's create a function and call the method getArea the parameter is IShape

// func calculateArea(shapeInp IShape) float64 {
// 	shp, ok := shapeInp.(SRectangle)

// 	if ok {
// 		println("The shape is a rectangle")
// 		return shp.getArea()
// 	} else {
// 		return 0
// 	}
// }

// func main() {
// 	// create a rectangle
// 	rect := SRectangle{length: 5, width: 10}
// 	// call the calculateArea function and pass the rectangle
// 	println(calculateArea(rect))
// }

package main

import (
	"errors"
	"fmt"
)

func divideByZero(val int) (float64, error) {
	if val == 0 {
		return 0, errors.New("Can't divide by 0")
	}
	return 1 / float64(val), nil
}

func main() {
	num, erMsg := divideByZero(0)
	if erMsg != nil {
		fmt.Println(erMsg)
	}
	fmt.Println(num)
}
