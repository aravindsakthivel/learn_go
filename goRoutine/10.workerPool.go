package main

import (
	"log"
	"sync"
	"time"
)

// Task definition

type Task struct {
	ID int
}

// Way to process the tasks
func (t *Task) Process() {
	log.Println("Processing task ", t.ID)
	// Simulate a time consuming process
	time.Sleep(2 * time.Second)
}

// Worker pool definition
type WorkerPool struct {
	Tasks       []Task
	concurrency int
	tasksChan   chan Task
	wg          sync.WaitGroup
}

// Function to execute worker pool

func (wp *WorkerPool) worker() {
	for task := range wp.tasksChan {
		task.Process()
		wp.wg.Done()
	}
}

func (wp *WorkerPool) Run() {
	// Initialize the task channel
	log.Println(len(wp.Tasks))
	wp.tasksChan = make(chan Task, len(wp.Tasks))

	// Start worker
	for i := 0; i < wp.concurrency; i++ {
		go wp.worker()
	}

	// Send tasks to the tasks Channel
	wp.wg.Add(len(wp.Tasks))
	for _, task := range wp.Tasks {
		// log.Println("Tasks ", task)
		wp.tasksChan <- task
	}
	close(wp.tasksChan)
	wp.wg.Wait()

}

func singleTask() {
	// Create new tasks
	tasks := make([]Task, 20)
	for i := 0; i < 20; i++ {
		tasks[i] = Task{ID: i + 1}
	}

	// Create a workerPool
	wp := WorkerPool{
		Tasks:       tasks,
		concurrency: 5, // number of workers that can run at a time
	}

	// Run the pool
	wp.Run()
	log.Println("All tasks are done")
}

// func main() {
// 	singleTask()
// }
