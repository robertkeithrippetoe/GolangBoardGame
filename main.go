package main

import (
	"fmt"
	"strconv"

	"github.com/GolangBoardGame/character"
)

func main() {
	fmt.Println("Now Booting Game")
	character1 := character.GenerateCharacter()
	fmt.Println("My name is " + character1.Name + "! My stats are: Endurance: " + strconv.Itoa(character1.Endurance) + " Health: " + strconv.Itoa(character1.Health) + " Speed: " + strconv.Itoa(character1.Speed) + " Strength " + strconv.Itoa(character1.Strength))
	character2 := character.GenerateCharacter()
	fmt.Println("My name is " + character2.Name + "! My stats are: Endurance: " + strconv.Itoa(character2.Endurance) + " Health: " + strconv.Itoa(character2.Health) + " Speed: " + strconv.Itoa(character2.Speed) + " Strength " + strconv.Itoa(character2.Strength))
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
