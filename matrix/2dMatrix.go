package main

import "fmt"

func createMatrix() {
	var matrix = [][]int{}

	for i := 0; i < 3; i++ {
		var row = []int{}
		for j := 0; j < 3; j++ {
			row = append(row, i+j)
		}
		matrix = append(matrix, row)
	}

	fmt.Println(matrix)
}

func searchWithRange() {
	var arr [5]int = [5]int{1, 2, 3, 4, 5}

	for ind, val := range arr {
		fmt.Println(ind, val)
	}
}

func main() {
	createMatrix()
	searchWithRange()
}
