package mechanics

import "fmt"

const (
	Physical ActionType = "Physical"
	Magical  ActionType = "Magical"
)

type ClassAbility interface {
	UseClassAbility(*Character) bool
	SetClassAbility(int, int)
	GetClassName() string
	RestoreAbility()
}

type Class struct {
	className         string
	defaultAttackType ActionType
	ClassAbilityUsed  int
	maxAbilityUses    int
}

func (class *Class) SetClassName(name string) {
	class.className = name
}

func (class *Class) SetAttackType(actionType ActionType) {
	class.defaultAttackType = actionType
}

type Archer struct {
	Class
	unicAbility DamageOverTime
}

func (archer *Archer) GetClassName() string {
	return archer.className
}

func (archer *Archer) SetClassAbility(duration int, value int) {
	archer.unicAbility = DamageOverTime{
		baseEffect: baseEffect{
			instantDamage:     0,
			instantDamageType: Magical,
			ActionType:        Magical,
			duration:          duration,
		},
		dotDamage: value,
	}
}

func (archer *Archer) UseClassAbility(target *Character) bool {
	if archer.ClassAbilityUsed < archer.maxAbilityUses {
		archer.unicAbility.ApplyEffect(target)
		archer.ClassAbilityUsed++
		fmt.Printf("%d/%d time(s) fire arraws was used!\n", archer.ClassAbilityUsed, archer.maxAbilityUses)
		fmt.Printf("%s will burn for %d turn(s) for %d damage\n", target.Name, archer.unicAbility.duration, archer.unicAbility.dotDamage)
		return true
	} else {
		return false
	}
}

func (archer *Archer) RestoreAbility() {
	archer.ClassAbilityUsed = 0
}

type Mage struct {
	Class
	unicAbility Freeze
}

func (mage *Mage) RestoreAbility() {
	mage.ClassAbilityUsed = 0
}

func (mage *Mage) GetClassName() string {
	return mage.className
}

func (mage *Mage) SetClassAbility(duration int, value int) {
	mage.unicAbility = Freeze{
		baseEffect: baseEffect{
			instantDamage:     value,
			instantDamageType: Magical,
			ActionType:        Magical,
			duration:          duration,
		},
		canAttack: false,
	}
}

func (mage *Mage) UseClassAbility(target *Character) bool {
	if mage.ClassAbilityUsed < mage.maxAbilityUses {
		mage.unicAbility.ApplyEffect(target)
		mage.ClassAbilityUsed++
		fmt.Printf("%d/%d time(s) freeze was used!\n", mage.ClassAbilityUsed, mage.maxAbilityUses)
		fmt.Printf("%s will not hit for %d turn(s) and take %d damage instantly\n", target.Name, mage.unicAbility.duration, mage.unicAbility.instantDamage)
		return true
	} else {
		return false
	}
}

type Knight struct {
	Class
	unicAbility Resistance
}

func (knight *Knight) GetClassName() string {
	return knight.className
}

func (knight *Knight) SetClassAbility(duration int, value int) {
	knight.unicAbility = Resistance{
		baseEffect: baseEffect{
			instantDamage:     0,
			instantDamageType: Magical,
			ActionType:        Physical, // type of reducing damage
			duration:          duration,
		},
		damageResist: float32(value) / 10.0,
	}
}

func (knight *Knight) RestoreAbility() {
	knight.ClassAbilityUsed = 0
}

func (knight *Knight) UseClassAbility(target *Character) bool {
	if knight.ClassAbilityUsed < knight.maxAbilityUses {
		knight.unicAbility.ApplyEffect(target)
		knight.ClassAbilityUsed++
		fmt.Printf("%d/%d times knight was shielded up\n", knight.ClassAbilityUsed, knight.maxAbilityUses)
		fmt.Printf("Hero will reduce %d persent of damage for %d turn(s)\n", int(knight.unicAbility.damageResist*100), knight.unicAbility.duration)
		return true
	} else {
		return false
	}

}

func CreateClass(class string) ClassAbility {
	switch class {
	case "Archer":
		return &Archer{
			Class: Class{
				className:         class,
				defaultAttackType: Physical,
				ClassAbilityUsed:  0,
				maxAbilityUses:    2,
			},
		}
	case "Mage":
		return &Mage{
			Class: Class{
				className:         class,
				defaultAttackType: Physical,
				ClassAbilityUsed:  0,
				maxAbilityUses:    1,
			},
		}
	case "Knight":
		return &Knight{
			Class: Class{
				className:         class,
				defaultAttackType: Physical,
				ClassAbilityUsed:  0,
				maxAbilityUses:    1,
			},
		}
	default:
		return nil
	}
}
