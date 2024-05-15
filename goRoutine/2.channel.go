package main

import "log"

func ChannelMain() {
	myChnl := make(chan int)

	go func() {
		myChnl <- 2
		myChnl <- 1
	}()

	resp := <-myChnl

	log.Println(resp)

}

// func main() {
// 	ChannelMain()
// }
