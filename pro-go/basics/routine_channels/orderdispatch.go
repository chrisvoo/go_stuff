package main

import (
	"fmt"
	"math/rand"
	"time"
)

type DispatchNotification struct {
	Customer string
	*Product
	Quantity int
}

var Customers = []string{"Alice", "Bob", "Charlie", "Dora"}

/*
The DispatchOrders function creates a random number of DispatchNotification values and sends
them through the channel that is received through the channel parameter.
There is no way to know in advance how many DispatchNotification values the DispatchOrders function will create.
To avoid to cause deadlock with all the goroutines blocker, the sender indicates when no further values are coming
through the channel, which is done by closing the channel.
The location of the arrow specifies the direction of the channel. When the arrow follows the chan keyword,
then the channel can be used only to send. The channel can be used to receive only if the arrow precedes the chan
keyword (<-chan, for example). Go allows bidirectional channels to be assigned to unidirectional channel variables,
allowing restrictions to be applied
*/
func DispatchOrders(channel chan<- DispatchNotification) {
	orderCount := rand.Intn(5) + 2
	fmt.Println("Order count:", orderCount)
	for i := 0; i < orderCount; i++ {
		channel <- DispatchNotification{
			Customer: Customers[rand.Intn(len(Customers)-1)],
			Quantity: rand.Intn(10),
			Product:  ProductList[rand.Intn(len(ProductList)-1)],
		}
		time.Sleep(time.Millisecond * 750)
	}
	// You need to close channels only when it is helpful to do so to coordinate your goroutines.
	// Go doesn’t require channels to be closed to free up resources or perform any kind of housekeeping task.
	close(channel)
}

func receiveDispatches(channel <-chan DispatchNotification) {
	for details := range channel {
		fmt.Println("Dispatch to", details.Customer, ":", details.Quantity,
			"x", details.Product.Name)
	}
	fmt.Println("Channel has been closed")
}

func enumerateProducts(channel chan<- *Product) {
	for _, p := range ProductList[:3] {
		channel <- p
		time.Sleep(time.Millisecond * 800)
	}
	close(channel)
}

func listProductsMultiChannel(channel1, channel2 chan<- *Product) {
	/* without default, the select statement will block until one of the channels can receive a value.
	   The distribution of the values is unpredictable and can be uneven
	*/
	for _, p := range ProductList {
		select {
		case channel1 <- p:
			fmt.Println("Send via channel 1")
		case channel2 <- p:
			fmt.Println("Send via channel 2")
		}
	}
	close(channel1)
	close(channel2)
}

/* A select statement can also be used to send to a channel without blocking */
func listProducts(channel chan<- *Product) {
	for _, p := range ProductList {
		select {
		case channel <- p:
			fmt.Println("Sent product:", p.Name)
		default:
			fmt.Println("Discarding product:", p.Name)
			time.Sleep(time.Second)
		}
	}
	close(channel)
}
