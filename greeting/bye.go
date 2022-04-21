package greeting

import (
	"errors"
	"fmt"
	"log"
)

func Bye(name string) (string, error) {
	log.SetPrefix("bye Error: ")
	if name == "" {
		log.Fatal(errors.New("you forgot a name in the bye function"))
		return "", nil
	}

	message := fmt.Sprintf("Bye, %v. Bye-Bye!", name)
	return message, nil
}
