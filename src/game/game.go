package game

import (
	"fmt"
	"math/rand"
	"rpg-saga/src/mechanics"
)

// TODO: турнирная сетка, боёвка, логирование
var classes = [...]string{"Archer", "Knight", "Mage"}

type Game struct {
	NumMembers       int
	ClasesAvailable  []mechanics.ClassAbility
	EffectsAvailable []mechanics.Effect
	Characters       []mechanics.Character
}

func (game *Game) InitCharacters() {
	Characters := []mechanics.Character{}
	for i := 0; i < game.NumMembers; i++ {
		name := "name" + fmt.Sprint(i+1)
		h := rand.Intn(51) + 50
		dmg := rand.Intn(10) + 10
		act := mechanics.CharacterAction{
			ActionType: mechanics.Physical,
			Damage:     dmg,
		}
		cls_num := rand.Intn(len(classes))

		char := mechanics.CreateChar(name, h, act, classes[cls_num])
		Characters = append(Characters, char)
	}
	game.Characters = Characters
	// for _, c := range Characters {
	// 	spew.Dump(c)
	// 	fmt.Println("")
	// 	fmt.Println("")
	// }
}

func Fight(f1 mechanics.Character, f2 mechanics.Character) mechanics.Character {
	a := Adjudicator{
		Figther1: f1,
		Figther2: f2,
	}
	w := a.StartFight()
	return w
}

type Arena struct {
	members []mechanics.Character
}

func (arena *Arena) StartRound() {
	fmt.Println("-")
	fmt.Println("New round!")
	fmt.Println("")
	len_characters := len(arena.members)
	var newMembers []mechanics.Character
	if len_characters > 1 {
		for i := 0; i < len_characters-1; i += 2 {
			figther1 := arena.members[i]
			figther2 := arena.members[i+1]
			winner := Fight(figther1, figther2)
			fmt.Println("winner is : " + winner.Name)
			newMembers = append(newMembers, winner)
		}
		arena.members = newMembers
	}
}

func StartGame() {
	members := 8
	if members%2 == 0 {
		game := Game{
			NumMembers: members,
		}
		game.InitCharacters()
		arena := Arena{members: game.Characters}
		for isEnded := false; isEnded == false; {
			arena.StartRound()
			if len(arena.members) < 2 {
				isEnded = true
			}
		}
	}
}
