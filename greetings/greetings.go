package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

func Hello(name string) (string, error) {
	// if no name was given, return an error with a message
	if name == "" {
		return "", errors.New("empty name")
	}
	// return a greeting that embeds the name in a message
	message := fmt.Sprintf(randomFormate(), name)
	return message, nil
}

func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)
	// loop through the received slice of names calling the Hello function to get a message for each name.

	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		// in the map, associate the retrieved message with the name.
		messages[name] = message
	}
	return messages, nil

}
func randomFormate() string {
	strMessage := []string{
		"Hi, %v Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}
	return strMessage[rand.Intn(len(strMessage))]
}
