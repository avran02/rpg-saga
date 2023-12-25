package mechanics

import (
	"errors"
	"fmt"
)

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
	currentHealth := float32(c.GetHealth())
	dmg := float32(action.Damage)
	if len(c.Effects) > 0 {
		for _, e := range c.Effects {
			switch effect := e.(type) {
			case Resistance:
				if effect.duration > 0 && effect.ActionType == action.ActionType {
					resisted := dmg * (1 - effect.damageResist)
					lastHp := currentHealth - resisted
					c.SetHealth(int(lastHp))
					effect.duration--
					fmt.Printf("After resist %s takes %d dmg\n", c.Name, int(resisted))
				}
			case DamageOverTime:
				if effect.duration > 0 {
					lastHp := currentHealth - float32(effect.dotDamage)
					c.SetHealth(int(lastHp))
					effect.duration--
					fmt.Printf("Fire deals %d dmg. %s will burn for %d turns\n", effect.dotDamage, c.Name, effect.duration)
				}
			default:
				lastHp := int(currentHealth - dmg)
				c.SetHealth(lastHp)
			}
		}
	} else {
		lastHp := int(currentHealth - dmg)
		c.SetHealth(lastHp)
	}
}

func (c *Character) SetHealth(h int) {
	c.health = h
}

func (c Character) GetHealth() int {
	return c.health
}

func (c *Character) addEffect(effect Effect) {
	c.Effects = append(c.Effects, effect)
}

func (c *Character) MakeAction(target *Character) (bool, error) {
	if len(c.Effects) > 0 {
		for _, e := range c.Effects {
			switch effect := e.(type) {
			case Freeze:
				if effect.canAttack == true && effect.duration < 1 {
					target.GetAction(c.CharacterAction)
					effect.duration--
					return true, nil
				} else {
					fmt.Println(c.Name, "can`t attack beacouse he`s cold")
					return false, nil
				}
			default:
				target.GetAction(c.CharacterAction)
				return true, nil
			}
		}
	} else {
		target.GetAction(c.CharacterAction)
		return true, nil
	}
	return false, errors.New("Make action unexcepted problem")
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
