package main

import (
	"fmt"
	"log"
	"example.com/greetings"
)

func main() {
	// Set properties of the predefined logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file and line number
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// A slice of names
	names := []string{"Josue", "Julia", "Chris"}

	// Request greeting messages for the names
	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	// If no erro was returned, print the returned message to the console
	for _, value := range messages {
		fmt.Println(value)
	}
}
