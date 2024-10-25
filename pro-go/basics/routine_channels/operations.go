package main

import (
	"fmt"
	"time"
)

func CalcStoreTotal(data ProductData) {
	var storeTotal float64
	/*
		I have set the size of the buffer to 2, meaning that two senders will be able to send values through the
		channel without having to wait for them to be received.
	*/
	var channel chan float64 = make(chan float64, 3)
	for category, group := range data {
		// When the Go runtime encounters the go keyword, it creates a new goroutine.
		// These statements are executed concurrently
		go group.TotalPrice(category, channel)
	}

	time.Sleep(time.Second * 5)
	fmt.Println("-- Starting to receive from channel")

	for i := 0; i < len(data); i++ {
		fmt.Println(len(channel), cap(channel))
		fmt.Println("-- channel read pending", len(channel), "items in buffer, size", cap(channel))
		// Receiving from a channel is a blocking operation, meaning that execution will not continue until a value has been received
		value := <-channel
		fmt.Println("-- channel read complete", value)
		storeTotal += value
		time.Sleep(time.Second)
	}
	fmt.Println("Total:", ToCurrency(storeTotal))
}

/* You may see different results based on the order in which keys are retrieved from the map */
func (group ProductGroup) TotalPrice(category string, resultChannel chan float64) {
	var total float64
	for _, p := range group {
		//fmt.Println(category, "product:", p.Name)
		total += p.Price
	}
	fmt.Println(category, "channel sending", ToCurrency(total))
	resultChannel <- total
	fmt.Println(category, "channel send complete")
}
