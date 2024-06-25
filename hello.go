package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {
	message := greetings.Hello("Yousef")
	fmt.Println(message)
}