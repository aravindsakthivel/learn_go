package main

import "fmt"

func main() {
	fmt.Println(proPlan, basicPlan)

	var arr = [5]int{1, 2, 3, 4, 5}
	var slices = arr[1:3]
	fmt.Println(slices, arr)

	var dyArr = []int{1, 2, 3, 4, 5}
	fmt.Println(dyArr)
}

const (
	proPlan   = 1
	basicPlan = 2
)
