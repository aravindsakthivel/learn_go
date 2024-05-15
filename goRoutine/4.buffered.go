package main

import "log"

// func writeBufferedChannel(chanel chan<- int) {
// 	chars := []int{1, 2, 3}
// 	for inf, s := range chars {
// 		log.Println("Writing to buffered channel: ", s, " at index: ", inf)
// 		// select {
// 		// case chanel <- s:
// 		// }
// 		chanel <- s

// 	}

// 	close(chanel)
// }

func checkBufferedChannel() {
	buffChnl := make(chan int, 3)

	chars := []int{1, 2, 3}
	for inf, s := range chars {
		log.Println("Writing to buffered channel: ", s, " at index: ", inf)
		select {
		case buffChnl <- s:
		}
	}

	close(buffChnl)

	for rcBfChnl := range buffChnl {
		log.Println("Received from buffered channel: ", rcBfChnl)
	}
}

// func main() {
// 	checkBufferedChannel()
// }
