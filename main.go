package main

import (
	"fmt"
	"saga/src/game"
	// "reflect"
)

func main() {
	fmt.Println("Started")
	err := game.StartGame()
	if err != nil {
		fmt.Println("Error: ", err)
	}
}
