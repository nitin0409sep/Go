package main

import (
	"fmt"
	"time"
)

// Sending
func processNum(numChan chan int) {
	for num := range numChan {
		fmt.Println("processing num", num)
		time.Sleep(time.Second * 2)
	}

	// This not gonna work
	// for i := 0; i < <-numChan; i++ {
	// 	fmt.Println("processing num", i)
	// time.Sleep(time.Second * 2)
	// }

	// This will work
	// for {
	// 	num, ok := <-numChan
	// 	if !ok {
	// 		break
	// 	}
	// 	fmt.Println(num)
	// }
}

// Receive
func sum(res chan int, num1 int, num2 int) {

	result := num1 + num2

	// Store result in channel
	res <- result
}

// Go Routine Func
func print(done chan bool) {
	fmt.Println("processing......")

	defer func() {
		done <- true
	}()
}

// Email Queue Function
func emailSender(emailChan <-chan string, done chan<- bool) {
	for email := range emailChan {
		fmt.Println("Sending email to ", email)
		time.Sleep(time.Second)
	}

	defer func() {
		done <- true
	}()
}

// Communications b/w the go routines is handled by channels -> for eg if you want to get/send data from 1 go routines into another go routines
func main() {

	// messageChannel := make(chan string) // Channels are blocking -> until second side is not ready to accept it

	// messageChannel <- "ping" // Inserting Message in to Channel

	// var msg string = <-messageChannel // Receiving Data from Channel

	// fmt.Println("msg", msg)

	// numChan <- 4

	// time.Sleep(time.Second * 1)

	// Send
	// numChan := make(chan int) // Created a channel that will accept int value

	// go processNum(numChan) // started a func that will run on go routines --- Channels are already reference types in Go.

	// for {
	// 	numChan <- rand.Intn(100) // rand.Intn -> Generated Random number b/w a range that you provides
	// }

	// Receive
	// result := make(chan int)

	// go sum(result, 1, 2)

	// res := <-result // Reading and wrting in channels are blocking -> Its partially true as we have another type of channels i.e. buffer channels

	// fmt.Println(res)

	// Go routines and channel -> If you don't wanna use wait group then you can go with channels too
	// done := make(chan bool) //! These are called as unbuffered channels -> Sending and Receiving things are blocking

	// go print(done)

	// <-done // This will also block the main thread -> Here we are Receiving the data in channel

	//! Buffered Channel
	emailChan := make(chan string, 100) // 2nd parameter is size
	done := make(chan bool)

	go emailSender(emailChan, done)

	for i := range 5 {
		emailChan <- fmt.Sprintf("%d@gmail.com", i)
	}

	fmt.Println("done sending")

	close(emailChan)
	<-done

	//! Waana access  data on 2 different channels
	chan1 := make(chan string)
	chan2 := make(chan int)

	go func() {
		chan1 <- "Channel 1"
	}()

	go func() {
		chan2 <- 20
	}()

	// Receive Data from Multiple Channels
	for i := 0; i < 2; i++ {
		select {
		case chan1Val := <-chan1:
			fmt.Println("Received Data form chan1", chan1Val)
		case chan2Val := <-chan2:
			fmt.Println("Received Data form chan1", chan2Val)
		}
	}

}
