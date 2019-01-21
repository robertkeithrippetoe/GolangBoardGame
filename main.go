package main

import (
	"fmt"
	"strconv"

	draft "github.com/GolangBoardGame/Draft"
)

func main() {
	fmt.Println("Now Booting Game")
	character1 := draft.GenerateCharacter()
	fmt.Println("My name is " + character1.Name + "! My stats are: " + strconv.Itoa(character1.Endurance) + strconv.Itoa(character1.Health) + strconv.Itoa(character1.Speed) + strconv.Itoa(character1.Strength))
	character2 := draft.GenerateCharacter()
	fmt.Println("My name is " + character2.Name + "! My stats are: " + strconv.Itoa(character2.Endurance) + strconv.Itoa(character2.Health) + strconv.Itoa(character2.Speed) + strconv.Itoa(character2.Strength))
	//character1 = draft.GenerateCharacter
	//fmt.Println("Say hi to " + character1.name)
	/*
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
	*/
}
