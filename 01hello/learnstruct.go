package main

import "fmt"

func structLesson() {
	type car struct {
		brand string
		info  string
	}
	var myCar car = car{brand: "SLFord"}
	fmt.Println(myCar)
	fmt.Println(myCar.info)
}

func anonymousStruct() {
	myCar := struct {
		brand string
	}{
		brand: "AFord",
	}
	fmt.Println(myCar)
}

type rect struct {
	width  int
	height int
}

// area has a receiver of (r rect)
func (r rect) area() int {
	return r.width * r.height
}

var r = rect{
	width:  5,
	height: 10,
}
