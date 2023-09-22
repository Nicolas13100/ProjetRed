package game

import (
	"fmt"
	"math/rand"
	"time"
)

var gridSize = 5
var currentRoom Carte
var currentFloor = 1

func InitMap(name string, description string) Carte {
	fmt.Println(name, description)
	rand.Seed(time.Now().UnixNano())

	initializeGame()

	for {
		printMap()
		movePlayer()

		if currentRoom.hasExit && P1.x == currentRoom.x && P1.y == currentRoom.y {
			fmt.Println("You found stairs! Moving to a new map...")
			currentFloor++
			initializeNewMap()
		}
		if currentRoom.HasMonster && P1.x == currentRoom.x && P1.y == currentRoom.y {
			fmt.Println("You encountered a monster!")
			// Add logic to handle the encounter (e.g., fight or escape)
			Fight(&P1, &Monstre1)
		}
	}
}

func initializeGame() {
	currentRoom = Carte{x: rand.Intn(gridSize), y: rand.Intn(gridSize), hasExit: true, floor: currentFloor}
	addMonsters()
}

func initializeNewMap() {
	currentRoom = Carte{x: rand.Intn(gridSize), y: rand.Intn(gridSize), hasExit: true, floor: currentFloor}
	addMonsters()
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
	for _, monster := range currentRoom.Monsters {
		if P1.x == monster.x && P1.y == monster.y {
			encounterMonster()
		}
	}
}

func addMonsters() {
	monsterCount := rand.Intn(4) // Generate a random number of monsters (0 to 3)
	for i := 0; i < monsterCount; i++ {
		room := Carte{x: rand.Intn(gridSize), y: rand.Intn(gridSize)}
		// Ensure the room doesn't already have a monster or player
		for room.HasMonster || (room.x == P1.x && room.y == P1.y) {
			room.x = rand.Intn(gridSize)
			room.y = rand.Intn(gridSize)
		}
		currentRoom.Monsters = append(currentRoom.Monsters, Monstre{x: room.x, y: room.y})
	}
}
func encounterMonster() {
	var monster Monstre
	if currentFloor >= 1 && currentFloor <= 4 {
		monsters := []Monstre{Goblin}
		randomIndex := rand.Intn(len(monsters))
		monster = monsters[randomIndex]
	} else {
		monsters := []Monstre{Goblin, Orc, HamsterGeant}
		randomIndex := rand.Intn(len(monsters))
		monster = monsters[randomIndex]
	}
	fmt.Println("You encountered a monster! What do you want to do?")
	fmt.Println("1. Fight")
	fmt.Println("2. Flee")

	var choice int
	fmt.Scan(&choice)

	switch choice {
	case 1:
		// Call the fight function
		Fight(&P1, &monster)
	case 2:
		// Add logic for fleeing (e.g., move player to a random position)
		fmt.Println("You chose to flee!")
		// Optionally, you can implement logic to move the player to a random position
	default:
		fmt.Println("Invalid choice. You freeze in fear!")
	}
}
