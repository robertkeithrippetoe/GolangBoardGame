package character

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type Character struct {
	Name      string
	Speed     int
	Strength  int
	Health    int
	Endurance int
}

func GenerateCharacter() Character {
	rand.Seed(time.Now().UnixNano())
	var character Character
	var firstName string
	var lastName string
	r := rand.Intn(10) + 1
	character.Speed = r
	r = rand.Intn(10) + 1
	character.Strength = r
	r = rand.Intn(10) + 1
	character.Health = r
	r = rand.Intn(10) + 1
	character.Endurance = r
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
	character.Name = firstName + " " + lastName
	return character
}

func PrintCharacter(character Character) {
	fmt.Println("My name is " + character.Name + "! My stats are: Endurance: " + strconv.Itoa(character.Endurance) + " Health: " + strconv.Itoa(character.Health) + " Speed: " + strconv.Itoa(character.Speed) + " Strength " + strconv.Itoa(character.Strength))
}
