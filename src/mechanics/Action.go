package mechanics

type ActionType string

type CharacterAction struct{
	ActionType
	damage int
	Effects []Effect 
}
