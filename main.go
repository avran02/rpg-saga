package main

import (
	"fmt"
	// "math/rand"
	"saga/src/game"
)

func main() {
	fmt.Println("Started")
	err := game.StartGame()
	if err != nil {
		fmt.Println("Error: ", err)
	}
}
