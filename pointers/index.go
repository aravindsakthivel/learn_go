package main

import "fmt"

func testPointer() {
	x := 10 // a created value can't be assigned to a pointer
	y := &x // y is a pointer to x

	fmt.Println(x, *y)
	*y = 8
	fmt.Println(x, *y)
	*y = 9
	// remove y pointer to x
	y = nil
	fmt.Println(x, y)
	y = &x
	fmt.Println(x, *y)
}

type cars struct {
	brand string
}

func (c *cars) changeBrand(brand string) {
	c.brand = brand
}

func main() {
	testPointer()

	c := cars{brand: "Toyota"}
	fmt.Println(c)

	// pointer was not mentioned in the function but still used
	c.changeBrand("Honda")

	fmt.Println(c)
}
