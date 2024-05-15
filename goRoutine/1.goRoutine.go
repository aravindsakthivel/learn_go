package main

import (
	"log"
	"time"
)

func runGoRoutine(inp int) {
	log.Println("Running go routine with input: ", inp)
}

func goroutine() {
	go runGoRoutine(1)
	go runGoRoutine(2)
	go runGoRoutine(3)

	time.Sleep(1 * time.Second)

	log.Println("Main function")
}

// func main() {
// 	goroutine()
// }
