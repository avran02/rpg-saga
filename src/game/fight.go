package game

import (
	"fmt"
	"math/rand"
	"saga/src/mechanics"
)

type Adjudicator struct {
	Figther1    mechanics.Character
	Figther2    mechanics.Character
	CurrentTurn int
}

func (a Adjudicator) StartFight() (mechanics.Character, error) {
	a.CurrentTurn = rand.Intn(2)
	var active *mechanics.Character
	var passive *mechanics.Character
	var isUsed bool
	fmt.Println(a.Figther1.Name, a.Figther1.Class.GetClassName(), a.Figther1.Damage, "dmg and", a.Figther1.GetHealth(), "hp")
	fmt.Println("VERSUS")
	fmt.Println(a.Figther2.Name, a.Figther2.Class.GetClassName(), a.Figther2.Damage, "dmg and", a.Figther2.GetHealth(), "hp")
	fmt.Println("==============================")
	for {
		if a.CurrentTurn%2 == 0 {
			active = &a.Figther1
			passive = &a.Figther2
		} else {
			active = &a.Figther2
			passive = &a.Figther1
		}
		tryUseAbility := rand.Intn(3) // set a chanse of using ability
		a.CurrentTurn++
		fmt.Printf("\nTURN %d\n%s (%d) hp attack %s (%d) hp\n", a.CurrentTurn, active.Name, active.GetHealth(), passive.Name, passive.GetHealth())
		if tryUseAbility != 1 {
			isAttacked, err := active.MakeAction(passive)
			if err != nil {
				fmt.Println(err)
				return a.Figther1, err
			}
			if isAttacked {
				fmt.Println(active.Name, "deals", active.Damage, "to", passive.Name, passive.GetHealth(), "left")
			}
		} else {
			isUsed = active.Class.UseClassAbility(passive)
			if !isUsed {
				isAttacked, err := active.MakeAction(passive)
				if err != nil {
					return a.Figther1, err
				}
				if isAttacked {
					fmt.Println(active.Name, "deals", active.Damage, "to", passive.Name, passive.GetHealth(), "left")
				}
			}
		}
		h1 := a.Figther1.GetHealth()
		h2 := a.Figther2.GetHealth()
		if h1 < 1 || h2 < 1 {
			if h1 > h2 {
				return a.Figther1, nil
			} else {
				return a.Figther2, nil
			}
		}
	}
}
