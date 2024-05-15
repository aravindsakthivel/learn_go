package main

import (
	"log"
)

func writeChannel(chanel chan<- int) {
	for i := 1; i <= 5; i++ {
		chanel <- i
	}

	close(chanel)
}

func buildAndCheck() {
	var chanel chan int = make(chan int)

	go writeChannel(chanel)
	i := 0
	for indChn := range chanel {
		println(indChn)
		i++
		log.Println("Received from chanel: ", i)
	}
}

// func main() {
// 	buildAndCheck()
// }
