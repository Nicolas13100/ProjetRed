package game

import (
	"fmt"
	"math/rand"
	"time"
)

type room struct {
	x, y    int
	hasExit bool
}

var gridSize = 5
var currentRoom room

func InitMap(name string, description string) room {
	rand.Seed(time.Now().UnixNano())

	initializeGame()

	for {
		printMap()
		movePlayer()

		if currentRoom.hasExit && P1.x == currentRoom.x && P1.y == currentRoom.y {
			fmt.Println("You found stairs! Moving to a new map...")
			initializeNewMap()
		}
	}
}

func initializeGame() {
	currentRoom = room{x: rand.Intn(gridSize), y: rand.Intn(gridSize), hasExit: rand.Intn(4) == 0}
}

func initializeNewMap() {
	currentRoom = room{x: rand.Intn(gridSize), y: rand.Intn(gridSize), hasExit: rand.Intn(4) == 0}
	P1.x = rand.Intn(gridSize)
	P1.y = rand.Intn(gridSize)
}

func printMap() {
	for i := 0; i < gridSize; i++ {
		for j := 0; j < gridSize; j++ {
			if i == P1.y && j == P1.x {
				fmt.Print("P ")
			} else if i == currentRoom.y && j == currentRoom.x && currentRoom.hasExit {
				fmt.Print("E ")
			} else {
				fmt.Print(". ")
			}
		}
		fmt.Println()
	}
}

func movePlayer() {
	var direction int
	fmt.Print("Enter direction (1.Nord/2.Ouest/3.Sud/4.Est): ")
	fmt.Scan(&direction)

	switch direction {
	case 1:
		if P1.y > 0 {
			P1.y--
		}
	case 2:
		if P1.x > 0 {
			P1.x--
		}
	case 3:
		if P1.y < gridSize-1 {
			P1.y++
		}
	case 4:
		if P1.x < gridSize-1 {
			P1.x++
		}
	}
}
