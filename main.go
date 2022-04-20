package main

import (
	"app/greeting"
	"fmt"

	"rsc.io/quote"
)

func main() {
	fmt.Println("Hello bro")
	fmt.Println(quote.Go())
	fmt.Println(greeting.Hello("Chris"))
	fmt.Println(greeting.Bye("Chris"))
}
