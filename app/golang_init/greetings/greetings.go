package greetings

import "fmt"

func Goodmornig (name string) string {
	message := fmt.Sprintf("Good morning, %s!", name)
	return message
}