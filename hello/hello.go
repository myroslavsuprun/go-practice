package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	// Logger predefining.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// Get a greeting message.
	message, err := greetings.Hello("Myroslav")

	// Error catch.
	if err != nil {
		log.Fatal(err)
	}

	// If there is no error occurance, just print the message.

	fmt.Println(message)
}
