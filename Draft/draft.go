package draft

import (
	"math/rand"
)

type Character struct {
	Name      string
	Speed     int
	Strength  int
	Health    int
	Endurance int
}

func GenerateCharacter() Character {
	var character Character
	var firstName string
	var lastName string
	r := rand.Intn(10) + 1
	character.speed = r
	r = rand.Intn(10) + 1
	character.strength = r
	r = rand.Intn(10) + 1
	character.health = r
	r = rand.Intn(10) + 1
	character.endurance = r
	r = rand.Intn(5) + 1
	switch r {
	case 1:
		firstName = "Holly"
	case 2:
		firstName = "Keith"
	case 3:
		firstName = "Mike"
	case 4:
		firstName = "Jasper"
	case 5:
		firstName = "Alex"
	}
	r = rand.Intn(5) + 1
	switch r {
	case 1:
		lastName = "Christofferson"
	case 2:
		lastName = "Brackenwald"
	case 3:
		lastName = "Chimboovist"
	case 4:
		lastName = "Swimtoodler"
	case 5:
		lastName = "Jistastimbolt"
	}
	character.name = firstName + " " + lastName
	return character
}
