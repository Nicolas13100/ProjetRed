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
		if len(currentRoom.Exits) == 0 {
			fmt.Println("Vous êtes bloqués, il n'y a aucune sortie de ce côté.")
			continue
		}
		fmt.Print("Vous avez trouvé la sortie : ")
		for direction := range currentRoom.Exits {
			fmt.Printf("%s ", direction)
		}
		fmt.Println()

		var chosenDirection string
		fmt.Print("Choisissez une direction : nord, sud, est, ouest ")
		fmt.Scanln(&chosenDirection)

		// Vérifiez si la direction choisie est une sortie valide
		nextRoom, ok := currentRoom.Exits[chosenDirection]
		if !ok {
			fmt.Println("Invalid direction. Try again.")
			continue
		}

		// Déplacez le joueur vers la nouvelle salle
		currentRoom = nextRoom
	}
}
