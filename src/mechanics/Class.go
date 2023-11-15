package mechanics 
// import (
//     "fmt"
// )
    

type ClassAbility interface {
    UseClassAbility(target *Character) bool
}

type Class struct{
    className string
    defaultAttackType ActionType
    ClassAbilityUsed int
    maxAbilityUses int
}


type Archer struct{
    Class
    unicAbility DamageOverTime
}
func (archer Archer) UseClassAbility(target *Character) bool{
    if archer.ClassAbilityUsed < archer.maxAbilityUses{
        archer.unicAbility.ApplyEffect(target)
        archer.ClassAbilityUsed++
        return true
    }else{
        return false
    }
} 


type Mage struct{
    Class
    unicAbility Freeze
}
func (mage Mage) UseClassAbility(target *Character) bool{
    if mage.ClassAbilityUsed < mage.maxAbilityUses{
        mage.unicAbility.ApplyEffect(target)
        mage.ClassAbilityUsed++
        return true
    }else{
        return false
    }
} 


type Knight struct{
    Class
    unicAbility Resistance
}
func (knight Knight) UseClassAbility(target *Character) bool{
    if knight.ClassAbilityUsed < knight.maxAbilityUses{
        knight.unicAbility.ApplyEffect(target)
        knight.ClassAbilityUsed++
        return true
    }else{
        return false
    }

} 