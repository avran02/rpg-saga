package game

import (
	"fmt"
	"math/rand"
	"rpg-saga/src/mechanics"

	"strconv"
)

type Adjudicator struct {
	Figther1    mechanics.Character
	Figther2    mechanics.Character
	CurrentTurn int
}

func (a *Adjudicator) StartFight() mechanics.Character {
	a.CurrentTurn = rand.Intn(2)
	var active mechanics.Character
	var passive mechanics.Character
	for {
		if a.CurrentTurn%2 == 0 {
			active = a.Figther1
			passive = a.Figther2
		} else {
			active = a.Figther2
			passive = a.Figther1
		}
		active.MakeAction(passive)
		fmt.Println(active.Name + "deals " + strconv.Itoa(active.Damage) + " to " + passive.Name)
		a.CurrentTurn++
		h1 := a.Figther1.GetHealth()
		h2 := a.Figther2.GetHealth()

		if h1 < 1 || h2 < 1 {
			return a.Figther1
		}
	}
}
