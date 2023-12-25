package mechanics

type Effect interface {
	ApplyEffect(c *Character)
}

type baseEffect struct {
	instantDamage     int
	instantDamageType ActionType
	ActionType        // type of DoT damage
	duration          int
}

type Resistance struct {
	damageResist float32
	baseEffect
}

func (effect Resistance) ApplyEffect(c *Character) {
	c.addEffect(effect)
}

type DamageOverTime struct {
	baseEffect
	dotDamage int
}

func (effect DamageOverTime) ApplyEffect(c *Character) {
	c.addEffect(effect)
}

type Freeze struct {
	baseEffect
	canAttack bool
}

func (effect Freeze) ApplyEffect(c *Character) {
	c.addEffect(effect)
}
