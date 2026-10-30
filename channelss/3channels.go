package main

import (
	"fmt"
	"time"
)

// receive only channel   <- direction use to secure only send channel
func emailSender(emailChan <-chan string, done chan<- bool) {
	defer func() { done <- true }()
	emailChan <- "pravin@gmail.com" //we cannot send
	for email := range emailChan {
		fmt.Println("Sending Email", email)
		time.Sleep(time.Second)
	}
}

func main() {

	emailChan := make(chan string, 100)

	done := make(chan bool)

	go emailSender(emailChan, done)

	for i := 1; i <= 5; i++ {

		emailChan <- fmt.Sprintf("%dexample@gmail.com", i)
	}
	fmt.Println("done email sending")
	close(emailChan)
	<-done
}
