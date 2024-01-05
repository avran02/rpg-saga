package mechanics

import (
	"fmt"
	"reflect"
)

var ResistancePointerType = reflect.TypeOf(&Resistance{})
var FreezePointerType = reflect.TypeOf(&Freeze{})
var DotPointerType = reflect.TypeOf(&DamageOverTime{})

type Character struct {
	Name    string
	hpPool  int
	health  int
	Effects []Effect
	CharacterAction
	Class ClassAbility
}

func (c *Character) RestoreChar() {
	c.health = c.hpPool
	c.Effects = []Effect{}
	c.Class.RestoreAbility()
}

func (c *Character) GetAction(action CharacterAction) {
	currentHealth := c.GetHealth()
	dmg := action.Damage
	var toDeal int
	var hasResist bool
	hasResist = false
	if len(c.Effects) > 0 {
		for _, e := range c.Effects {
			et := reflect.ValueOf(e).Type()
			if et == ResistancePointerType {
				if e.GetDuration() > 0 && e.GetActionType() == action.ActionType {
					resisted := float32(dmg) * (float32((100 - e.GetValue())) / 100) // тут совсем херня
					toDeal += int(resisted)
					hasResist = true
					e.SetDuration(e.GetDuration() - 1)
					fmt.Printf("After resist %s takes %d dmg\n", c.Name, int(resisted))
				}
			} else if et == DotPointerType {
				if e.GetDuration() > 0 {
					toDeal += e.GetValue()
					e.SetDuration(e.GetDuration() - 1)
					fmt.Printf("Fire deals %d dmg. %s will burn for %d turns\n", e.GetValue(), c.Name, e.GetDuration())
				} else {
					fmt.Println("fire with no dmg( dur:", e.GetDuration())
				}
			} else {
				fmt.Println("continue")
				continue
			}
		}
	}
	if hasResist {
		c.SetHealth(currentHealth - toDeal)
	} else {
		c.SetHealth(currentHealth - toDeal - int(dmg))
	}
}

func (c *Character) SetHealth(h int) {
	c.health = h
}

func (c Character) GetHealth() int {
	return c.health
}

func (c *Character) addEffect(effect Effect) {
	effect.DealInstDmg(c)
	c.Effects = append(c.Effects, effect)
}

func (c *Character) MakeAction(target *Character) (bool, error) { // if has freeze, ok also false
	if len(c.Effects) > 0 {
		for _, e := range c.Effects {
			et := reflect.ValueOf(e).Type()
			if et == FreezePointerType {
				if e.GetDuration() > 0 {
					fmt.Println(c.Name, "can`t attack beacouse he`s cold")
					e.SetDuration(e.GetDuration() - 1)
					return false, nil
				} else {
					target.GetAction(c.CharacterAction)
					return true, nil
				}
			} else {
				continue
			}
		}
	}
	target.GetAction(c.CharacterAction)
	return true, nil
}

func CreateChar(n string, h int, act CharacterAction, cslName string) Character {
	char := Character{
		Name:            n,
		hpPool:          h,
		CharacterAction: act,
		Class:           CreateClass(cslName),
	}
	char.SetHealth(h)
	return char
}
