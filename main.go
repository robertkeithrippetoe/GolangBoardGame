package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/GolangBoardGame/character"
	"github.com/GolangBoardGame/player"
)

func main() {
	fmt.Println("Now Booting Game")
	numCharacters := 5
	roster := make([]character.Character, numCharacters)
	for i := 0; i < numCharacters; i++ {
		tempCharacter := character.GenerateCharacter()
		roster[i] = tempCharacter
		character.PrintCharacter(roster[i])
		time.Sleep(5) // Ensures characters don't get repeated if they are generated too fast
	}

	numPlayers := 2
	playerArray := make([]player.Player, numPlayers)
	var player1 player.Player
	player1.Name = "player1"
	playerArray[0] = player1
	var player2 player.Player
	player2.Name = "player2"
	playerArray[1] = player2

	fmt.Println("Pick your characters!")

	for i := 0; i < len(roster); i++ {
		fmt.Print(i + 1)
		fmt.Print(". ")
		character.PrintCharacter(roster[i])
	}

	fmt.Println(player1.Name + " which character do you choose?")
	var reply string
	fmt.Scanln(&reply)
	replyInt, _ := strconv.ParseInt(reply, 10, 64)
	player1.Team = append(player1.Team, roster[replyInt-1])
	character.PrintCharacter(player1.Team[0])
	roster = append(roster[:replyInt-1], roster[replyInt:]...)

	for i := 0; i < len(roster); i++ {
		fmt.Print(i + 1)
		fmt.Print(". ")
		character.PrintCharacter(roster[i])
	}

	fmt.Println(player2.Name + " which character do you choose?")
	fmt.Scanln(&reply)
	replyInt, _ = strconv.ParseInt(reply, 10, 64)
	player2.Team = append(player2.Team, roster[replyInt-1])
	character.PrintCharacter(player2.Team[0])
	playerArray[0] = player1
	playerArray[1] = player2
	player.AllPrint(playerArray)
	/*
		for i := 0; i < len(roster); i++ {
			fmt.Print(i + 1)
			fmt.Print(". ")
			character.PrintCharacter(roster[i])
		}

		fmt.Println(player2.Name + " which character do you choose?")
		fmt.Scanln(&reply)
		replyInt, _ = strconv.ParseInt(reply, 10, 64)
		player2.Team[0] = roster[replyInt]
	*/

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
