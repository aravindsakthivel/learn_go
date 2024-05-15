package main

import (
	"fmt"
	"sync"
)

func mapMain() {
	var wg sync.WaitGroup
	var m sync.Map

	// Function to store values concurrently
	storeValue := func(key, value interface{}) {
		defer wg.Done()
		m.Store(key, value)
	}

	// Function to load and print values concurrently
	loadValue := func(key interface{}) {
		defer wg.Done()
		if value, ok := m.Load(key); ok {
			fmt.Printf("Key: %v, Value: %v\n", key, value)
		} else {
			fmt.Printf("Key: %v not found\n", key)
		}
	}

	// Storing values concurrently
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go storeValue(i, fmt.Sprintf("Value %d", i))
	}

	// Loading values concurrently
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go loadValue(i)
	}

	// Wait for all goroutines to finish
	wg.Wait()
}

// func main() {
// 	mapMain()
// }
