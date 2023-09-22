package game

import "fmt"

func NewRoom(name, description string, hasMonster bool) *Room {
	return &Room{
		Name:        name,
		Description: description,
		Exits:       make(map[string]*Room),
		HasMonster:  hasMonster,
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

func Dongeon(P1 *Personnage) {
	forest := NewRoom("Forêt", "Vous arrivé dans une forêt dense et mystérieuse", true)
	entrance := NewRoom("Entrée", "Vous êtes à l'entrée du dongeon", false)
	hall := NewRoom("Hall", "Vous êtes dans un long hall", true)
	treasureRoom := NewRoom("Salle au trésor", "Vous êtes dans une pièce remplie de trésors.", false)

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
		fmt.Print("Vous avez trouvé la sortie : ")
		for direction := range currentRoom.Exits {
			fmt.Printf("%s ", direction)
		}
		fmt.Println()

		var chosenDirection int
		fmt.Print("Choisissez une direction : (1 : Nord, 2 : Sud, 3 : Est, 4 : ouest ")
		fmt.Scanln(&chosenDirection)
		var direction string
		switch chosenDirection {
		case 1:
			direction = "nord"
		case 2:
			direction = "sud"
		case 3:
			direction = "est"
		case 4:
			direction = "ouest"
			if chosenDirection < 1 || chosenDirection > 4 {
				fmt.Println("Direction invalide. Veuillez réessayer.")
				continue
			}
		}

		// Vérifiez si la direction choisie est une sortie valide
		nextRoom, ok := currentRoom.Exits[direction]
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
