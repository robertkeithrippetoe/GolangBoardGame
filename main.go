package main

import (
	"fmt"
)

func main() {
	fmt.Println("What is your name?")
	var text string
	fmt.Scanln(&text)
	switch text {
	case "Keith":
		fmt.Println("Wow, do you listen to internet radio!?")
	case "Mike":
		fmt.Println("Wow, do you know about K&M!?")
	default:
		fmt.Println("I don't care.")
	}
}
