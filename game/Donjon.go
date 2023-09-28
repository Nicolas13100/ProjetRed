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
			fmt.Println("Vous avez trouvé les escaliers permettant de monter, passage à l'étage supérieur.")
			currentFloor++
			initializeNewMap()
		} else if currentRoom.hasReturn && P1.x == stairsX && P1.y == stairsY && currentFloor > 1 {
			fmt.Println("Vous avez trouvé les escaliers permettant de redescendre, retour à l'étage inférieur.")
			currentFloor--
			initializeNewMap()
		} else if currentRoom.hasReturn && P1.x == stairsX && P1.y == stairsY && currentFloor <= 1 {
			var choix int
			fmt.Println("Vous êtes au rez-de-chaussée, souhaitez-vous sortir du donjon ? (1 : Oui / 2 : Non)")
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
	for _, jar := range currentRoom.Jars {
		if P1.x == jar.x && P1.y == jar.y {
			var loot = Item{
				Name: "Fourrure de Loup",
			}
			encounterJar(P1, loot)
		}
	}
	for _, monster := range currentRoom.Monsters {
		if P1.x == monster.x && P1.y == monster.y {
			encounterMonster(Monstre1)
		}
	}
}

func initializeGame() {
	currentRoom = Carte{x: rand.Intn(gridSize), y: rand.Intn(gridSize), hasExit: true, hasReturn: true, HasJar: true, floor: currentFloor}
	addMonsters()
	addJar()
}

func initializeNewMap() {
	currentRoom = Carte{x: rand.Intn(gridSize), y: rand.Intn(gridSize), hasExit: true, hasReturn: true, HasJar: true, floor: currentFloor}
	addMonsters()
	addJar()
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

func addJar() {
	JarCount := rand.Intn(4) // Generate a random number of jars (0 to 5)
	for i := 0; i < JarCount; i++ {
		room := Carte{x: rand.Intn(gridSize), y: rand.Intn(gridSize)}
		// Ensure the room doesn't already have a jar or player
		for room.HasJar || (room.x == P1.x && room.y == P1.y) {
			room.x = rand.Intn(gridSize)
			room.y = rand.Intn(gridSize)
		}
		currentRoom.Jars = append(currentRoom.Jars, Jar{x: room.x, y: room.y})
	}
}

func newJar() *Jar {
	jar := &Jar{}
	randomContentType := ContentType(rand.Intn(3))
	jar.ContentType = randomContentType

	switch randomContentType {
	case Loot:
		jar.HasLoot = true
		jar.Contenu = []Item{
			{Name: "Fourrure de Loup", Quantity: 1},
			{Name: "Epée de l'aventurier", Quantity: 1},
		}
	case Monster:
		jar.HasMonster = true
	case Event:
		jar.HasEvent = true
	}
	return jar
}

func (j *Jar) openJar() Item {
	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(len(j.Contenu))
	item := j.Contenu[index]
	return item
}

func encounterJar(P1 Personnage, item Item) {

	rand.Seed(time.Now().UnixNano())
	jar := newJar()

	fmt.Println("Vous avez trouvé une jar ! Qu'allez-vous faire?")
	fmt.Println("1. Ouvrir")
	fmt.Println("2. Continuer votre route.")
	var choice int
	fmt.Scan(&choice)
	switch choice {
	case 1:
		if jar.HasLoot {
			loot := jar.openJar()
			fmt.Printf("Vous avez obtenu un(e) %s!\n", loot.Name)
			fmt.Println(P1.Inventory)
			P1.Inventory[loot.Name] += 1
		} else if jar.HasMonster {
			monsters := []Monstre{Jarpent}
			Monstre1 := monsters[0]
			fmt.Println("Vous êtes tombé sur un monstre !")
			Fight(&P1, &Monstre1)
		} else if jar.HasEvent {
			var Question int
			randomQuestion := rand.Intn(3)
			Question = randomQuestion
			switch Question {
			case 1:
				var choix string
				fmt.Println("Quel être, pourvu d'une seule voix, a d'abord quatre jambes le matin, puis deux jambes à midi, et trois jambes le soir ? ")
				fmt.Scan(&choix)
				if choix == "Homme" || choix == "homme" {
					fmt.Println("Comment ?!... Tu as trouvé la bonne réponse, voici ta récompense.")
					P1.Inventory["Collier du Sphinx"]++
					fmt.Println("Vous avez obtenu le Collier du Sphinx !")
					return
				} else {
					fmt.Println("Non, ce n'est pas la bonne réponse, déguste ma papatate")
					P1.Hp -= 25
					fmt.Println("Vous avez perdu 25 PV.")
					return
				}
			case 2:
				var choix string
				fmt.Println("On la reçoit sans la remercier. On en profite sans savoir pourquoi. On la donne à d'autres. On la perd sans s'en apercevoir. Qui suis-je ?")
				fmt.Scan(&choix)
				if choix == "La vie" || choix == "Vie" || choix == "vie" || choix == "la vie" {
					fmt.Println("Ouiii, c'est la bonne réponse, voici ta récompense.")
					P1.Inventory["La Clé"]++
					fmt.Println("Vous avez obtenu La Clé !")
					return
				} else {
					fmt.Println("Non, ce n'est pas la bonne réponse, la Boule t'amène à l'étage inférieur")
					currentFloor -= 1
					return
				}
			case 3:
				var choix string
				fmt.Println("On la reçoit sans la remercier. On en profite sans savoir pourquoi. On la donne à d'autres. On la perd sans s'en apercevoir. Qui suis-je ?")
				fmt.Scan(&choix)
				if choix == "La vie" || choix == "Vie" || choix == "vie" || choix == "la vie" {
					fmt.Println("Ouiii, c'est la bonne réponse, voici ta récompense.")
					P1.Inventory["La Clé"]++
					fmt.Println("Vous avez obtenu La Clé !")
					return
				} else {
					fmt.Println("Non, ce n'est pas la bonne réponse, la Boule t'amène à l'étage inférieur")
					currentFloor -= 1
					return
				}
			}
		} else {
			fmt.Println("La jar est vide.")
		}
	case 2:
		fmt.Println("Vous continuez votre route.")
	default:
		fmt.Println("Choix inconnu, vous êtes pris de peur !")
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

func encounterMonster(Monstre1 Monstre) {
	if currentFloor >= 1 && currentFloor <= 4 {
		monsters := []Monstre{Goblin}
		randomIndex := rand.Intn(len(monsters))
		Monstre1 = monsters[randomIndex]
	} else {
		monsters := []Monstre{Goblin, Orc, HamsterGeant}
		randomIndex := rand.Intn(len(monsters))
		Monstre1 = monsters[randomIndex]
	}
	fmt.Println("Vous rencontez un monstre ! Qu'allez-vous faire?")
	fmt.Println("1. Combatre")
	fmt.Println("2. Fuir")

	var choice int
	fmt.Scan(&choice)

	switch choice {
	case 1:
		Fight(&P1, &Monstre1)
	case 2:
		fmt.Println("Vous fuyez !")
		return

	default:
		fmt.Println("Choix inconnue, vous êtes pris de peur !")
	}
}

var (
	JarLoot = Jar{
		HasLoot:    true,
		HasMonster: false,
		HasEvent:   false,
	}
	JarMonster = Jar{
		HasLoot:    false,
		HasMonster: false,
		HasEvent:   false,
	}
	JarEvent = Jar{
		HasLoot:    false,
		HasMonster: false,
		HasEvent:   true,
	}
)
