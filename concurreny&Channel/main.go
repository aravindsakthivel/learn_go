package main

import (
	"fmt"
	"time"
)

func controller1() {
	// Create a buffered channel with a capacity of 3
	ch := make(chan int, 3)

	// Sender goroutine
	go func() {
		for i := 1; i <= 5; i++ {
			ch <- i
			fmt.Println("Sent:", i)
		}
		close(ch) // Close channel after sending all messages
	}()

	// Receiver goroutine
	go func() {
		for value := range ch {
			fmt.Println("Received:", value)
			time.Sleep(1 * time.Second) // simulate time to process
		}
	}()

	// Wait for user input to end the program (simulating waiting for all operations to complete)
	fmt.Scanln()
}

func printValues(ch <-chan int) {
	for value := range ch {
		fmt.Println("Received:", value)
	}
}

func controller2() {
	ch := make(chan int)

	go func() {
		for i := 1; i <= 5; i++ {
			fmt.Println("Closing channel")
			ch <- i
		}

		// close(ch)
	}()
	fmt.Println("Waiting for values")
	printValues(ch)
}

func sendValues(ch chan<- int) {
	for i := 1; i <= 5; i++ {
		ch <- i
		fmt.Println("Sent:", i)
		time.Sleep(1 * time.Second) // Simulating some work
	}
	close(ch)
}

func controller3() {
	ch := make(chan int)

	go sendValues(ch)
	fmt.Println("Waiting for values")
	for value := range ch {
		fmt.Println("Received:", value)
	}
	fmt.Println("Channel closed")
}

func main() {
	controller1()
	// controller2()
	// controller3()
	// buildAndCheck()
}
