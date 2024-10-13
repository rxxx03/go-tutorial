package greetings

import (
	"errors"
	"fmt"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("name is missing")
	}
	messsage := fmt.Sprintf("Hi, %v. Welcome!", name)
	return messsage, nil
}
