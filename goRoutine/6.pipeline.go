package main

import "log"

func sliceNumber(arr []int) <-chan int {
	out := make(chan int, 3)

	go func() {
		for _, n := range arr {
			out <- n
			log.Println("Writing to channel: 11 ", n)
		}
		close(out)
	}()

	return out
}

func squareNumber(inp <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		for n := range inp {
			out <- n * n
			log.Println("Writing to channel: 25 ", n*n)
		}
		close(out)
	}()

	return out
}

func pipeline() {
	nums := []int{1, 2, 3, 4, 5}

	// Create a function to slice the numbers

	dataChanel := sliceNumber(nums)

	// Create a function to square the numbers

	squareChanel := squareNumber(dataChanel)

	// Print the squared numbers

	for n := range squareChanel {
		log.Println(n)
	}
}

// func main() {
// 	pipeline()
// }
