//Code for Two channels

package main

import "fmt"

func main() {
	chan1 := make(chan int)
	chan2 := make(chan string)
	go func() {
		chan1 <- 9767
	}()

	go func() {
		chan2 <- "Pravin Bendke"
	}()
	for i := 0; i < 2; i++ {
		select {
		case chan1Val := <-chan1:
			fmt.Println("receiving value of channel 1", chan1Val)
		case chan2Val := <-chan2:
			fmt.Println("receiving value of channel 2", chan2Val)
		}
	}
}
