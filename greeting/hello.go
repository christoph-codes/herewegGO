package greeting

import (
	"errors"
	"fmt"
	"log"
)

func Hello(name string) (string, error) {
	log.SetPrefix("helloError: ")
	if name == "" {
		log.Fatal(errors.New("you forgot a name in the hello function"))
		return "", nil
	}

	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil
}
