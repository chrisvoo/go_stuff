package main

import (
	"fmt"
)

// initialize the channel that will be receiving the messages
var messages = make(chan string)

func pong() {
	go createPing()
	msg := <-messages
	fmt.Println(msg)
}

func createPing() {
	messages <- "ping"
}

/* passing data from one Goroutine to another Goroutine via channels */
