package main

import (
	"fmt"
	"runtime"
)

func coresFn() {
	// Get the current GOMAXPROCS value
	currentValue := runtime.GOMAXPROCS(0)

	fmt.Printf("Current GOMAXPROCS value: %d\n", currentValue)

	// Set GOMAXPROCS to utilize all available CPU cores
	maxCores := runtime.NumCPU()
	fmt.Printf("maxCores value: %d\n", maxCores)
	newValue := maxCores - 2

	runtime.GOMAXPROCS(newValue)
	fmt.Printf("Updated GOMAXPROCS value: %d\n", runtime.GOMAXPROCS(0))
}

// func main() {
// 	coresFn()
// }
