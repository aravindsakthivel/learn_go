package main

import (
	"log"
)

func routine1(chl chan int) {
	chl <- 1
}

func routine2(chl chan int) {
	chl <- 2
}

func SelectMain() {

	ch1 := make(chan int)
	ch2 := make(chan int)

	go routine1(ch1)
	go routine2(ch2)

	select {
	case res := <-ch1:
		log.Println("Received from ch1: ", res)
	case res1 := <-ch2:
		log.Println("Received from ch2: ", res1)
	}
}

// func main() {
// 	SelectMain()
// }
