package game

import (
	"errors"
	"fmt"
	"math/rand"
	"saga/src/mechanics"
)

// TODO: турнирная сетка, боёвка, логирование, тесты
var classes = []string{"Archer", "Knight", "Mage"}

var heroNames = [...]string{
	"Aiden", "Bella", "Caden", "Daisy", "Elijah", "Freya", "Gavin", "Hazel",
	"Isaac", "Jasmine", "Kaden", "Lily", "Mason", "Nora", "Oliver", "Penelope",
	"Quinn", "Riley", "Samuel", "Sophia", "Tristan", "Uma", "Vincent", "Willow",
	"Xander", "Yara", "Zane",
}

type Game struct {
	NumMembers       int
	ClasesAvailable  []string
	EffectsAvailable []mechanics.Effect
	Characters       []mechanics.Character
}

func (game *Game) InitCharacters() {
	Characters := []mechanics.Character{}
	for i := 0; i < game.NumMembers; i++ {
		numNames := len(heroNames)
		randomIdx := rand.Intn(numNames)
		name := heroNames[randomIdx]
		h := rand.Intn(51) + 50
		dmg := rand.Intn(10) + 10
		act := mechanics.CharacterAction{
			ActionType: mechanics.Physical,
			Damage:     dmg,
		}
		cls_num := rand.Intn(len(classes))
		char := mechanics.CreateChar(name, h, act, classes[cls_num])
		abilityDuration := rand.Intn(2) + 1
		abilityValue := rand.Intn(6) + 1
		char.Class.SetClassAbility(abilityDuration, abilityValue)
		Characters = append(Characters, char)
	}
	game.Characters = Characters
}

func Fight(f1 mechanics.Character, f2 mechanics.Character) (mechanics.Character, error) {
	a := Adjudicator{
		Figther1: f1,
		Figther2: f2,
	}
	w, err := a.StartFight()
	if err != nil {
		return w, err
	}
	return w, nil
}

type Arena struct {
	members []mechanics.Character
}

func (arena *Arena) StartRound() error {
	fmt.Println("New round!")
	fmt.Println("")
	len_characters := len(arena.members)
	var newMembers []mechanics.Character
	if len_characters > 1 {
		for i := 0; i < len_characters-1; i += 2 {
			figther1 := arena.members[i]
			figther2 := arena.members[i+1]
			winner, err := Fight(figther1, figther2)
			if err != nil {
				return err
			}
			fmt.Println("winner is : " + winner.Name)
			fmt.Println("He has", winner.GetHealth(), "hp \n\n==============================")
			winner.RestoreChar()
			newMembers = append(newMembers, winner)
		}
		arena.members = newMembers
	}
	return nil
}

func StartGame() error {
	members := 8
	if members%2 == 0 {

		game := Game{
			NumMembers:      members,
			ClasesAvailable: classes,
		}
		game.InitCharacters()
		arena := Arena{members: game.Characters}
		for {
			arena.StartRound()
			if len(arena.members) < 2 {
				break
			}
		}
	} else {
		return errors.New("Invalid number of members")
	}
	return nil
}
