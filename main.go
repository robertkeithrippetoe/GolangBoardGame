package main

import (
	"fmt"

	"github.com/GolangBoardGame/character"
)

func main() {
	fmt.Println("Now Booting Game")
	character1 := character.GenerateCharacter()
	character.PrintCharacter(character1)
	character2 := character.GenerateCharacter()
	character.PrintCharacter(character2)
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
