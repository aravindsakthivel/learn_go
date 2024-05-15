package main

import (
	"log"
	"sync"
	"time"
)

var lock sync.Mutex

func process(data int) int {
	time.Sleep(time.Second * 2)
	return data * 2
}

func processData(wg *sync.WaitGroup, result *[]int, data int) {

	defer wg.Done()
	processedData := process(data)
	lock.Lock()
	*result = append(*result, processedData)
	lock.Unlock()
}

func processConfinement(wg *sync.WaitGroup, result *int, data int) {
	defer wg.Done()

	processedData := process(data)

	*result = processedData

}

func mutexesMain() {
	start := time.Now()
	var wg sync.WaitGroup

	input := []int{1, 2, 3, 4, 5}
	result := make([]int, len(input))

	for i, data := range input {
		wg.Add(1)
		go processConfinement(&wg, &result[i], data)
	}

	wg.Wait()
	log.Println(time.Since(start))
	log.Println(result)
	log.Println("Process done")

}

// func main() {
// 	mutexesMain()
// }
