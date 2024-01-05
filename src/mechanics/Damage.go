package mechanics

type Effect interface {
	ApplyEffect(c *Character)
	DealInstDmg(c *Character)
	GetDuration() int
	SetDuration(new int)
	GetActionType() ActionType
	GetValue() int
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

func (effect *Resistance) ApplyEffect(c *Character) {
	c.addEffect(effect)
}

func (e *Resistance) DealInstDmg(c *Character) {
	c.SetHealth(c.GetHealth() - e.instantDamage)
}

func (e *Resistance) GetActionType() ActionType {
	return e.ActionType
}

func (e *Resistance) GetDuration() int {
	return e.duration
}

// returns percent of resisting damage
func (e *Resistance) GetValue() int {
	return int(e.damageResist * 100)
}

func (e *Resistance) SetDuration(new int) {
	e.duration = new
}

type DamageOverTime struct {
	baseEffect
	dotDamage int
}

func (effect *DamageOverTime) ApplyEffect(c *Character) {
	c.addEffect(effect)
}

func (e *DamageOverTime) DealInstDmg(c *Character) {
	c.SetHealth(c.GetHealth() - e.instantDamage)
}

func (e *DamageOverTime) GetActionType() ActionType {
	return e.ActionType
}

func (e *DamageOverTime) GetDuration() int {
	return e.duration
}

// returns dot damage
func (e *DamageOverTime) GetValue() int {
	return int(e.dotDamage)
}

func (e *DamageOverTime) SetDuration(new int) {
	e.duration = new
}

type Freeze struct {
	baseEffect
	canAttack bool
}

func (effect *Freeze) ApplyEffect(c *Character) {
	c.addEffect(effect)
}

func (e *Freeze) DealInstDmg(c *Character) {
	c.SetHealth(c.GetHealth() - e.instantDamage)
}

func (e *Freeze) GetActionType() ActionType {
	return e.ActionType
}

func (e *Freeze) GetDuration() int {
	return e.duration
}

// returns instant damage
func (e *Freeze) GetValue() int {
	return int(e.instantDamage)
}

func (e *Freeze) SetDuration(new int) {
	e.duration = new
}
