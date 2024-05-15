package main

import (
	"log"
	"sync"
	"time"
)

type MultiTask interface {
	process()
}

// Build a email task definition

type EmailTask struct {
	Email       string
	Subject     string
	MessageBody string
}

// Create a execution function for the task

func (e *EmailTask) process() {
	log.Println("Processing task ", e.Email)
	time.Sleep(2 * time.Second)
}

// Build a Image recognition definition

type ImageRecognitionTask struct {
	imageUrl string
}

// Create a execution function for the task

func (i *ImageRecognitionTask) process() {
	log.Println("Processing task ", i.imageUrl)
	time.Sleep(5 * time.Second)
}

// Build a worker pool definition

type WorkersPool struct {
	Tasks       []MultiTask
	concurrency int
	taskChan    chan MultiTask
	wg          sync.WaitGroup
}

// Create a function to execute the tasks execute function

func (wp *WorkersPool) Worker() {
	for task := range wp.taskChan {
		task.process()
		wp.wg.Done()
	}
}

// Create a function for generation of worker pool and trigger the task

func (wp *WorkersPool) Run() {
	wp.taskChan = make(chan MultiTask, len(wp.Tasks))

	for i := 0; i < wp.concurrency; i++ {
		go wp.Worker()
	}

	wp.wg.Add(len(wp.Tasks))

	for _, task := range wp.Tasks {
		wp.taskChan <- task
	}

	close(wp.taskChan)

	wp.wg.Wait()
}

func multiTask() {
	tasks := []MultiTask{
		&EmailTask{Email: "email@code1", Subject: "test", MessageBody: "test"},
		&ImageRecognitionTask{imageUrl: "url1"},
		&EmailTask{Email: "email@code2", Subject: "test", MessageBody: "test"},
		&ImageRecognitionTask{imageUrl: "url2"},
		&EmailTask{Email: "email@code3", Subject: "test", MessageBody: "test"},
		&ImageRecognitionTask{imageUrl: "url3"},
		&EmailTask{Email: "email@code4", Subject: "test", MessageBody: "test"},
		&ImageRecognitionTask{imageUrl: "url4"},
		&EmailTask{Email: "email@code5", Subject: "test", MessageBody: "test"},
		&ImageRecognitionTask{imageUrl: "url5"},
		&EmailTask{Email: "email@code6", Subject: "test", MessageBody: "test"},
		&ImageRecognitionTask{imageUrl: "url6"},
		&EmailTask{Email: "email@code7", Subject: "test", MessageBody: "test"},
		&ImageRecognitionTask{imageUrl: "url7"},
		&EmailTask{Email: "email@code8", Subject: "test", MessageBody: "test"},
		&ImageRecognitionTask{imageUrl: "url8"},
		&EmailTask{Email: "email@code9", Subject: "test", MessageBody: "test"},
		&ImageRecognitionTask{imageUrl: "url9"},
		&EmailTask{Email: "email@code10", Subject: "test", MessageBody: "test"},
		&ImageRecognitionTask{imageUrl: "url10"},
	}

	wp := WorkersPool{
		Tasks:       tasks,
		concurrency: 5,
	}

	wp.Run()

	log.Println("All tasks are done")
}

// func main() {
// 	var wg sync.WaitGroup
// 	log.Println(wg)
// 	multiTask()
// }
