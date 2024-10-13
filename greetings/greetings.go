package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("name is missing")
	}
	messsage := fmt.Sprintf(randomFormat(), name)
	return messsage, nil
}

func randomFormat() string {
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Cor blimey, guvner. 'ow's me old china %v?",
	}

	return formats[rand.Intn(len(formats))]
}

