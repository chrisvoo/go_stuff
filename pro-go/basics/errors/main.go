package main

import (
	"fmt"
	"os"
)

func main() {
	arg := "sync"
	if len(os.Args) > 1 {
		arg = os.Args[1]
	}

	/* uses the defer keyword to register a function, which will be executed
	when the main function has completed, even if there has been no panic */
	defer func() {
		if arg := recover(); arg != nil {
			if err, ok := arg.(error); ok {
				fmt.Println("Error from recover function: ", err.Error())
			} else if str, ok := arg.(string); ok {
				fmt.Println("Message: ", str)
			} else {
				fmt.Println("Panic recovered")
			}
		}
	}()

	categories := []string{"Watersports", "Chess", "Running"}

	switch arg {
	case "sync":
		{
			for _, cat := range categories {
				total, err := Products.TotalPrice(cat)
				if err != nil {
					fmt.Println(err.Error())
				} else {
					fmt.Println(cat, "Total:", ToCurrency(total))
				}
			}
		}
	case "async":
		{
			var prodChannel chan ChannelMessage = make(chan ChannelMessage)
			go Products.TotalPriceAsync(categories, prodChannel)
			for item := range prodChannel {
				if item.CategoryError != nil {
					//fmt.Println(item.CategoryError)
					panic(item.CategoryError)
				} else {
					fmt.Printf("Total for %s: %s\n", item.Category, ToCurrency(item.Total))
				}
			}
		}
	case "panic":
		{
			channel := make(chan CategoryCountMessage)
			go processCategories(categories, channel)
			for message := range channel {
				if (message.TerminalError == nil) {
					fmt.Println(message.Category, "Total:", message.Count)
				} else {
					fmt.Println("An error occurred.")
				}
			}
		}

	default:
		fmt.Println("Wrong arg", arg)
	}
}
