package greetings

import "fmt"

func Hello(name string) string {

	// Returns a greeting that embeds the name in a message

	// shorthand variable declaration and initilization. Only valid inside a function body
	messsage := fmt.Sprintf("Hi, %v. Welcome!", name)

	// In go an error results if a variable is not used
	return messsage
}
