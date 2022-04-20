package greeting

import "fmt"

func Bye(name string) string {
	fmt.Println("Hi")
	message := fmt.Sprintf("Bye, %v. Bye-Bye!", name)
	return message
}
