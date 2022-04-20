package main

import (
	"app/greeting"
	"fmt"

	"rsc.io/quote"
)

func main() {
	fmt.Println("Hello bro")
	fmt.Println(quote.Glass())
	fmt.Println(greeting.Bye("Chris"))
}
