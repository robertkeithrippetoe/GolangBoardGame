package player

import (
	"fmt"

	character "github.com/GolangBoardGame/character"
)

// A Player in the game
type Player struct {
	Name     string
	TeamName string
	Team     []character.Character
}

// GeneratePlayer makes a player"
func GeneratePlayer() {

}

func AllPrint(playerArray []Player) {
	fmt.Println("These are all characters on the board.")
	for _, player := range playerArray {
		for _, characters := range player.Team {
			character.PrintCharacter(characters)
		}
	}
}
