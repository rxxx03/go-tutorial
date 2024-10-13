package greetings

import "fmt"

func Hello(name string) string {
	messsage := fmt.Sprintf("Hi, %v. Welcome!", name)
	return messsage
}