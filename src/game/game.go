package game
import(
	"rpg-saga/src/mechanics"
	"fmt"
)
// type Character struct{
// 	name string
// 	health int
// 	Effects []Effect
// 	CharacterAction
// 	Class ClassAbility
// }
type Game struct{
	NumMembers int
	clasesAvailable []mechanics.ClassAbility
	effectsAvailable []mechanics.Effect
}
func (game Game) InitGame(){
	Characters := []Character{}
	for i:=0 ; i< game.NumMembers ; i++{
		char := mechanics.Character{
			Name: "name" + fmt.Sprint(i+1),
		}
		Characters = append(Characters, char)
		// fmt.Println(char.Name)
	}
	fmt.Println(Characters)
}
func (game Game) StartGame(){}