package main

import "fmt"
import "rsc.io/quote"
import "golang_alpine/golang_init/greetings"

func main () {
		fmt.Println("Hello, World!")
		fmt.Println(quote.Go())
		message := greetings.Goodmornig("Shohei")
		fmt.Println(message)
}
