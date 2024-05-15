package main

import (
	"fmt"
	"time"
)

func sendMessage(msg message) string {
	return msg.getMessage()
}

type message interface {
	getMessage() string
}

// don't edit below this line

type birthdayMessage struct {
	birthdayTime  time.Time
	recipientName string
}

func (bm birthdayMessage) getMessage() string {
	return fmt.Sprintf("Hi %s, it is your birthday on %s", bm.recipientName, bm.birthdayTime.Format(time.RFC3339))
}

type sendingReport struct {
	reportName    string
	numberOfSends int
}

func (sr sendingReport) getMessage() string {
	return fmt.Sprintf(`Your "%s" report is ready. You've sent %v messages.`, sr.reportName, sr.numberOfSends)
}

func mainIn() {
	bdMessage := birthdayMessage{birthdayTime: time.Now(), recipientName: "John"}
	sdReport := sendingReport{reportName: "Birthday", numberOfSends: 10}
	// The birthdayMessage struct implements the message interface, because it has a method getMessage() string it will implicitly implement the message interface
	fmt.Println(sendMessage(bdMessage))
	fmt.Println(sendMessage(sdReport))
}
