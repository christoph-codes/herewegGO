package random

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"time"
)

func RandomFormat(name string) (string, error) {
	log.SetPrefix("Random Format: ")
	rand.Seed(time.Now().UnixNano())

	if name == "" {
		log.Fatal(errors.New("you forgot a name in the randomformat function"))
		return "", nil
	}

	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	// Return a randomly selected message format by specifying
	// a random index for the slice of formats.
	message := fmt.Sprintf(formats[rand.Intn(len(formats))], name)

	return message, nil
	// return formats[rand.Intn(len(formats))]
}
