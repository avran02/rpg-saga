package mechanics
// import (
// 	"fmt"
// )

type Character struct{
	Name string
	health int
	Effects []Effect
	CharacterAction
	Class ClassAbility
}

func (c Character) GetAction(action CharacterAction) {
	current := c.GetHealth()
	var dmg float32
	for _, e := range c.Effects{
		switch effect:= e.(type) {
		case Resistance:
			if effect.ActionType == action.ActionType{
				dmg = float32(c.damage) * (1.0 - float32(action.damage))
			}
		}
	}
	c.setHealth(int(float32(current) - dmg))
}

func (c *Character) setHealth(h int) {
    c.health = h
}

func (c Character) GetHealth() int {
    return c.health
}

func (c *Character) addEffect(effect Effect) {
    c.Effects = append(c.Effects, effect)
}

func (c Character) MakeAction(target Character){
	for _, e := range c.Effects{
		switch effect:= e.(type) {
		case Freeze:
			if effect.canAttack == true{
				target.GetAction(c.CharacterAction)
			}
		}
	}
}