package greeting

import "fmt"

func Hello(name string) string {
	fmt.Println("Hi")
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}
