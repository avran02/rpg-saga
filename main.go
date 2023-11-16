package main

import (
	// "rpg-saga/src/game"
	"fmt"
	"math/rand"
)

heroNames := []string{
	"Aiden", "Bella", "Caden", "Daisy", "Elijah", "Freya", "Gavin", "Hazel",
	"Isaac", "Jasmine", "Kaden", "Lily", "Mason", "Nora", "Oliver", "Penelope",
	"Quinn", "Riley", "Samuel", "Sophia", "Tristan", "Uma", "Vincent", "Willow",
	"Xander", "Yara", "Zane",
}

func main() {
	for i := 1; i < 20; i++ {
		a := rand.Intn(2)
		fmt.Println(a)
	}
	// game.StartGame()
}
