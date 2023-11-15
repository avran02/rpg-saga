package main
import (
	"rpg-saga/src/game"
	// "rpg-saga/src/mechanics"
	// "fmt"
)


func main(){
	members := 8
	if members % 2 == 0{
		gameInst := game.Game{
			NumMembers: members,
		}
		gameInst.InitGame()
	}
}
