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

	// names for greetings
	names := []string{"Myroslav", "James", "Michael"}

	messages, err := greetings.Hellos(names)

	// Error catch.
	if err != nil {
		log.Fatal(err)
	}

	// If there is no error occurance, just print the message.

	fmt.Println(messages)
}
