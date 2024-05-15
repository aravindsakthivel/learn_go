package main

import (
	"fmt"
	"sync"
	"time"
)

type SafeCounter struct {
	mu    sync.RWMutex
	value int
}

// Increment the counter with write lock
func (c *SafeCounter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
	fmt.Println("Incremented to:", c.value)
}

// Get the counter value with read lock
func (c *SafeCounter) Get() int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.value
}

func rwMain() {
	counter := SafeCounter{}

	var wg sync.WaitGroup

	// Increment the counter in several goroutines
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Increment()
		}()
	}

	// Read the counter value in several goroutines
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			time.Sleep(100 * time.Millisecond) // Ensure some increments happen first
			fmt.Printf("Goroutine %d: Counter value: %d\n", id, counter.Get())
		}(i)
	}

	wg.Wait()
	fmt.Println("Final counter value:", counter.Get())
}

func main() {
	rwMain()
}
