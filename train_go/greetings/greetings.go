package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}
	message := fmt.Sprintf(randomFormart(), name)
	return message, nil
}

func randomFormart() string {
	formart := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	return formart[rand.Intn(len(formart))]
}

func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		messages[name] = message
	}
	return messages, nil
}
