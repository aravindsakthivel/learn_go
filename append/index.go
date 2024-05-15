package main

import "fmt"

func main() {
	var array []int = []int{1, 2, 3, 4, 5}
	array = append(array, 6)
	fmt.Println(array)
}
