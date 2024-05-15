package main

import (
	"log"
	"time"
)

func doneFn(done <-chan bool) {
	for {
		select {
		case <-done:
			clg := <-done
			log.Println("Done: ", clg)
			return
		default:
			log.Println("Working...")
		}
	}
}

func DoneMain() {
	done := make(chan bool)
	go doneFn(done)

	time.Sleep(1 * time.Second)

	close(done)
}

// func main() {
// 	DoneMain()
// }
