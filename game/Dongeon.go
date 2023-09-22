package game

import (
	"fmt"
	"math/rand"
	"time"
)

func NewRoom(name, description string, exits map[string]*Room) *Room {
	return &Room{
		Name:        name,
		Description: description,
		Exits:       exits,
		HasMonster:  randomBool(),
	}
}

func (r *Room) AddExit(direction string, room *Room) {
	r.Exits[direction] = room
	room.Exits[reverseDirection(direction)] = r
}

func reverseDirection(direction string) string {
	switch direction {
	case "north":
		return "south"
	case "south":
		return "north"
	case "east":
		return "west"
	case "west":
		return "east"
	default:
		return ""
	}
}

func Dungeon(P1 *Personnage) {
	forest := NewRoom("Forêt", "Vous arrivez dans une forêt dense et mystérieuse", nil)
	entrance := NewRoom("Entrée", "Vous êtes à l'entrée du donjon", nil)
	hall := NewRoom("Hall", "Vous êtes dans un long hall", nil)
	treasureRoom := NewRoom("Salle au trésor", "Vous êtes dans une pièce remplie de trésors.", nil)
	directions := []string{"nord", "sud", "est", "ouest"}
	currentRoom := forest
	monsterRooms := map[*Room]bool{
		forest:       true,
		entrance:     false,
		hall:         true,
		treasureRoom: false,
	}
	for {
		fmt.Println(currentRoom.Description)
		if monster, ok := monsterRooms[currentRoom]; ok && monster {
			fmt.Println("Un monstre est apparu, préparez vous au combat !")
			Fight(P1, &Goblin)
		}

		if len(currentRoom.Exits) == 0 {
			fmt.Println("Vous êtes bloqués, il n'y a aucune sortie de ce côté.")
			continue
		}
		fmt.Print("Vous avez trouvé la sortie : ")
		for direction := range currentRoom.Exits {
			fmt.Printf("%s ", direction)
		}
		fmt.Println()

		var chosenDirection int
		fmt.Print("Choisissez une direction : 1.Nord, 2.Sud, 3.Est, 4.Ouest ")
		_, err := fmt.Scanln(&chosenDirection)

		if err != nil {
			fmt.Println("Invalid input. Please enter a number.")
			continue
		}

		// Check if the chosen direction is valid
		nextRoom, ok := currentRoom.Exits[directions[chosenDirection-1]]
		if !ok {
			fmt.Println("Invalid direction. Try again.")
			continue
		}

		// Move the player to the new room
		currentRoom = nextRoom
		if !ok {
			fmt.Println("Invalid direction. Try again.")
			continue
		}
		if nextRoom == nil {
			fmt.Println("La sortie est bloquée. Choisissez une autre direction.")
			continue
		}
		// Déplacez le joueur vers la nouvelle salle
		currentRoom = nextRoom
	}
}

func randomBool() bool {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(2) == 1
}
