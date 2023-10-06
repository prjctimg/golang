package main

import (
	"fmt"
	"golang/greetings"
)

func main() {
	//Get a greeting message and print it
	message := greetings.Hello("Dean")
	fmt.Println(message)
}
