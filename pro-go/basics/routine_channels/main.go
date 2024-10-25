package main

import (
	"fmt"
	"os"
	"time"
)

func main() {

	arg := "r"
	if len(os.Args) > 1 {
		arg = os.Args[1]
	}

	switch arg {
	case "r":
		{
			fmt.Println("main function started")
			CalcStoreTotal(Products)
			// time.Sleep(time.Second * 5) // avoid the main to finish before the goroutines terminates their job
			fmt.Println("main function complete")
		}
	case "c":
		{
			dispatchChannel := make(chan DispatchNotification, 100)
			var sendOnlyChannel chan<- DispatchNotification = dispatchChannel
			// var receiveOnlyChannel <-chan DispatchNotification = dispatchChannel
			go DispatchOrders(sendOnlyChannel)

			productChannel := make(chan *Product)
			go enumerateProducts(productChannel)
			openChannels := 2
			//receiveDispatches(receiveOnlyChannel)
			// receiveDispatches((<-chan DispatchNotification)(dispatchChannel))
			for {
				/*
					A select statement has a similar structure to a switch statement, except that the case statements
					are channel operations. When the select statement is executed, each channel operation is evaluated
					until one that can be performed without blocking is reached
				*/
				select {
				case details, ok := <-dispatchChannel:
					if ok {
						fmt.Println("Dispatch to", details.Customer, ":",
							details.Quantity, "x", details.Product.Name)
					} else {
						fmt.Println("Channel has been closed")
						dispatchChannel = nil
						openChannels--
					}
				case product, ok := <-productChannel:
					if ok {
						fmt.Println("Product:", product.Name)
					} else {
						fmt.Println("Product channel has been closed")
						productChannel = nil
						openChannels--
					}
				/*
					If the default clause is omitted, then the select statement will block until one of the
					channels has a value to be received. This can be useful, but it does not deal with channels
					that can be closed.
				*/
				default:
					if openChannels == 0 {
						goto alldone
					}
					fmt.Println("-- No message ready to be received")
					time.Sleep(time.Millisecond * 500)
				}
			}
		alldone:
			fmt.Println("All values received")
		}
	case "s":
		{
			productChannel := make(chan *Product, 5)
			go listProducts(productChannel)
			time.Sleep(time.Second)
			for p := range productChannel {
				fmt.Println("Received product:", p.Name)
			}
		}
	case "multi":
		{
			c1 := make(chan *Product, 2)
			c2 := make(chan *Product, 2)
			go listProductsMultiChannel(c1, c2)
			time.Sleep(time.Second)
			for p := range c1 {
				fmt.Println("Channel 1 received product:", p.Name)
			}
			for p := range c2 {
				fmt.Println("Channel 2 received product:", p.Name)
			}
		}
	default: fmt.Println("Cazzo passi?")
	}
}
