package mechanics

type ActionType string

type CharacterAction struct {
	ActionType
	Damage  int
	Effects []Effect
}
