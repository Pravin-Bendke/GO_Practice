package main

import (
	"fmt"
	"time"
)

//receive
// func sum(result chan int, num1 int, num2 int) {
// 	numResult := num1 + num2
// 	result <- numResult
// }

// sending data
//	func processNum(nums chan int) {
//		for num := range nums {
//			fmt.Println("Process number ", num)
//			time.Sleep(time.Second)
//		}
//	}

func emailSender(emailChan <-chan string, done chan bool) {
	defer func() { done <- true }()

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

	//receive
	// result := make(chan int)
	// go sum(result, 10, 20
	// res := <-result
	// fmt.Println(res)

	// send
	// numChan := make(chan int)
	// go processNum(numChan)
	// for {
	// 	numChan <- rand.Intn(100)
	// }

	// numChan := make(chan int)
	// go processNum(numChan)
	// for {
	// 	numChan <- rand.Intn(100)
	// }
	// numChan <- 5
	// time.Sleep(time.Second * 2)
	// messageChan := make(chan string)
	// messageChan <- "Pravin Bendke"
	// msg := <-messageChan
	// fmt.Println(msg)
}
