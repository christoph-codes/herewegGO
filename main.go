package main

import (
	"app/greeting"
	"fmt"
	"log"

	"rsc.io/quote"
)

func main() {
	log.SetPrefix("mainError: ")
	fmt.Println("Hello bro")
	// Print what comes from quote package
	fmt.Println(quote.Go())
	// Print from the hello function
	fmt.Println(greeting.Hello("Christopher"))
	// Print from the bye function
	fmt.Println(greeting.Bye("Coder"))
	// Print from the hello function to kill the entire application after this
	fmt.Println(greeting.Hello(""))
	fmt.Println(greeting.Bye("Yo"))
}
