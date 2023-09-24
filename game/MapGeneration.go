package game

import (
	"fmt"
	"math/rand"
	"time"
)

var gridSize = 5
var currentRoom Carte
var currentFloor = 1
var stairsX int
var stairsY int

func InitMap(name string, description string) Carte {
	fmt.Println(name, description)
	rand.Seed(time.Now().UnixNano())

	initializeGame()

	for {
		printMap()
		movePlayer()

		if currentRoom.hasExit && P1.x == currentRoom.x && P1.y == currentRoom.y {
			fmt.Println("Vous avez trouvé les escaliers permettant de descendre, passage à l'étage inférieur.")
			currentFloor--
			initializeNewMap()
		}
		if currentRoom.hasReturn && P1.x == stairsX && P1.y == stairsY && currentFloor > 1 {
			fmt.Println("Vous avez trouvé les escaliers permettant de remonter, retour à l'étage supérieur.")
			currentFloor++
			initializeNewMap()
		} else if currentRoom.hasReturn && P1.x == stairsX && P1.y == stairsY && currentFloor <= 1 {
			var choix int
			fmt.Println("Vous êtes au premier étage, souhaitez-vous sortir du donjon ? (1 : Oui / 2 : Non)")
			fmt.Scan(&choix)
			switch choix {
			case 1:
				fmt.Println("Vous êtes sorti du donjon.")
				return currentRoom
			case 2:
				fmt.Println("Vous avez choisi de rester dans le donjon.")
				continue
			}
		}

		if currentRoom.HasMonster && P1.x == currentRoom.x && P1.y == currentRoom.y {
			fmt.Println("You encountered a monster!")
			// Add logic to handle the encounter (e.g., fight or escape)
			Fight(&P1, &Monstre1)
		}
	}
}

func initializeGame() {
	currentRoom = Carte{x: rand.Intn(gridSize), y: rand.Intn(gridSize), hasExit: true, hasReturn: true, floor: currentFloor}
	addMonsters()
}

func initializeNewMap() {
	currentRoom = Carte{x: rand.Intn(gridSize), y: rand.Intn(gridSize), hasExit: true, hasReturn: true, floor: currentFloor}
	addMonsters()
	P1.x = rand.Intn(gridSize)
	P1.y = rand.Intn(gridSize)
	currentRoom.hasReturn = currentRoom.hasExit
	stairsX = P1.x
	stairsY = P1.y
}

func printMap() {
	for i := 0; i < gridSize; i++ {
		for j := 0; j < gridSize; j++ {
			if i == P1.y && j == P1.x {
				fmt.Print("P ")
			} else if i == currentRoom.y && j == currentRoom.x && currentRoom.hasExit {
				fmt.Print("E ")
			} else if i == stairsY && j == stairsX {
				fmt.Print("R ")
			} else {
				fmt.Print(". ")
			}
		}
		fmt.Println()
	}
}

func movePlayer() {
	var direction int
	fmt.Print("Enter direction (5.Nord/1.Ouest/2.Sud/3.Est): ")
	fmt.Scan(&direction)

	switch direction {
	case 5:
		if P1.y > 0 {
			P1.y--
		}
	case 1:
		if P1.x > 0 {
			P1.x--
		}
	case 2:
		if P1.y < gridSize-1 {
			P1.y++
		}
	case 3:
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
	monsterCount := rand.Intn(6) // Generate a random number of monsters (0 to 5)
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
	fmt.Println("Vous rencontez un monstre ! Qu'allez-vous faire?")
	fmt.Println("1. Combatre")
	fmt.Println("2. Fuir")

	var choice int
	fmt.Scan(&choice)

	switch choice {
	case 1:

		Fight(&P1, &monster)

	case 2:

		fmt.Println("Vous fuyez !")
		return

	default:
		fmt.Println("Choix unconnue, vous êtes pris de peur !")
	}
}
