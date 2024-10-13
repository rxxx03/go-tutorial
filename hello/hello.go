package main

import (
	"fmt"
	"ralf/greetings"
)

func main() {
	message := greetings.Hello("Toni")
	fmt.Println(message)
}
