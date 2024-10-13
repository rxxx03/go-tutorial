package main

import (
	"fmt"
	"log"

	"ralf/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	message, err := greetings.Hello("Toni")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}
