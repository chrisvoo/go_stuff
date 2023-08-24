package main

import (
	"fmt"
	"greetings"
	"log"
)

func main() {
	// Get a greeting message and print it.
	message, err := greetings.Hello("Chris")
	if err != nil {
		// this stops the execution here
		log.Fatal(err)
	}
	fmt.Println(message)

	// A slice of names.
	names := []string{"Gladys", "Samantha", "Darrin"}

	// Request greeting messages for the names.
	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(messages)
}
