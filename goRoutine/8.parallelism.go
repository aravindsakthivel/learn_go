package main

import (
	"fmt"
	"runtime"
	"sync"
)

func printNumbers(prefix string, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 5; i++ {
		fmt.Println(prefix, i)
	}
}

func maxProcMain() {
	runtime.GOMAXPROCS(2) // Set the maximum number of CPUs that can execute simultaneously

	var wg sync.WaitGroup
	wg.Add(2)

	go printNumbers("Goroutine 1", &wg)
	go printNumbers("Goroutine 2", &wg)

	wg.Wait()

}

// func main() {
// 	maxProcMain()
// }
