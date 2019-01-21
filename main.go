package main

import (
	"fmt"

	"github.com/GolangBoardGame/character"
	"github.com/GolangBoardGame/player"
)

func main() {
	fmt.Println("Now Booting Game")
	var roster [5]character.Character
	var player1 player.Player
	player1.Name = "player1"
	var player2 player.Player
	player2.Name = "player2"
	character1 := character.GenerateCharacter()
	roster[0] = character1
	character.PrintCharacter(character1)
	character2 := character.GenerateCharacter()
	roster[1] = character2
	character.PrintCharacter(character2)
	character3 := character.GenerateCharacter()
	roster[2] = character3
	character.PrintCharacter(character3)
	character4 := character.GenerateCharacter()
	roster[3] = character4
	character.PrintCharacter(character4)
	character5 := character.GenerateCharacter()
	roster[4] = character5
	character.PrintCharacter(character5)

	//character.PrintCharacter(character5)
	for i := 0; i < 5; i++ {
		fmt.Println(i)
		character.PrintCharacter(roster[i])
	}
	fmt.Println(player1.Name + " which character do you choose?")
	var reply string
	fmt.Scanln(&reply)
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
