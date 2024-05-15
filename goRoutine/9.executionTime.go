// package main

// import (
// 	"fmt"
// 	"log"
// 	"time"
// )

// func printNumbers1() {
// 	for i := 1; i <= 10; i++ {
// 		fmt.Println(i)
// 		// time.Sleep(100 * time.Millisecond) // Simulate work
// 	}
// }

// func main() {
// 	start := time.Now()
// 	log.Println("Start time:", start)

// 	printNumbers1()

// 	end := time.Now()
// 	log.Println("End time:", end)
// 	log.Println("Duration:", end.Sub(start))
// 	fmt.Scanln()
// }

package main

import (
	"fmt"
	"time"
)

func printNumbers2() {
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
		time.Sleep(100 * time.Millisecond) // Simulate work
	}
}

// func main() {
// 	start := time.Now()
// 	log.Println("Start time:", start)

// 	done := make(chan bool)
// 	go func() {
// 		printNumbers2()
// 		done <- true
// 	}()

// 	<-done

// 	end := time.Now()
// 	log.Println("End time:", end)
// 	log.Println("Duration:", end.Sub(start))
// 	fmt.Scanln()
// }

// package main

// import (
// 	"fmt"
// 	"log"
// 	"sync"
// 	"time"
// )

// const workerCount = 3

// func worker(id int, tasks <-chan int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	for task := range tasks {
// 		fmt.Println(task)
// 		// time.Sleep(100 * time.Millisecond) // Simulate work
// 	}
// }

// func main() {
// 	start := time.Now()
// 	log.Println("Start time:", start)

// 	tasks := make(chan int, 10)
// 	var wg sync.WaitGroup

// 	// Start worker pool
// 	for i := 0; i < workerCount; i++ {
// 		wg.Add(1)
// 		go worker(i, tasks, &wg)
// 	}

// 	// Send tasks
// 	for i := 1; i <= 10; i++ {
// 		tasks <- i
// 	}
// 	close(tasks)

// 	wg.Wait()

// 	end := time.Now()
// 	log.Println("End time:", end)
// 	log.Println("Duration:", end.Sub(start))
// 	fmt.Scanln()
// }

// package main

// import (
// 	"fmt"
// 	"log"
// 	"runtime"
// 	"sync"
// 	"time"
// )

// const parallelWorkers = 2

// func parallelWorker(id int, tasks <-chan int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	for task := range tasks {
// 		fmt.Printf("Worker %d: %d\n", id, task)
// 		time.Sleep(100 * time.Millisecond) // Simulate work
// 	}
// }

// func main() {
// 	runtime.GOMAXPROCS(parallelWorkers) // Set the maximum number of CPUs to use

// 	start := time.Now()
// 	log.Println("Start time:", start)

// 	tasks := make(chan int, 10)
// 	var wg sync.WaitGroup

// 	// Start parallel workers
// 	for i := 0; i < parallelWorkers; i++ {
// 		wg.Add(1)
// 		go parallelWorker(i, tasks, &wg)
// 	}

// 	// Send tasks
// 	for i := 1; i <= 10; i++ {
// 		tasks <- i
// 	}
// 	close(tasks)

// 	wg.Wait()

// 	end := time.Now()
// 	log.Println("End time:", end)
// 	log.Println("Duration:", end.Sub(start))
// }
