package game

import (
	"fmt"
)

type Personnage struct {
	Name      string
	Race      string
	Level     int
	HpMax     int
	Hp        int
	Inventory map[string]int
	Skills    []string
	Gold      int
}

func InitPersonnage(name string, race string) *Personnage {
	p1 := &Personnage{
		Name:   name,
		Race:   race,
		Level:  1,
		HpMax:  100,
		Hp:     40,
		Gold:   100, // Assuming an initial gold amount for the Personnage
		Skills: []string{"coup de poing"},
		Inventory: map[string]int{
			"Potion de soin":   3,
			"Potion de poison": 3,
		},
	}
	return p1
}

func (p1 *Personnage) DisplayInfo() {
	// Affichage des informations du Personnage p1

	for {
		fmt.Println("Nom:", p1.Name)
		fmt.Println("Classe:", p1.Race)
		fmt.Println("Niveau:", p1.Level)
		fmt.Println("Points de vie maximum:", p1.HpMax)
		fmt.Println("Points de vie actuels:", p1.Hp)
		fmt.Println("Cash:", p1.Gold)
		// Affichage de l'inventaire
		fmt.Println("\nType 'return' to come back to main menu")

		// Read user input
		var input string
		fmt.Scan(&input)

		// Check if the user wants to return to the main menu
		if input == "return" {
			break
		}
	}
}

func (p1 *Personnage) AccessInventory() {
	for {
		fmt.Println("\nInventaire:")
		for key, value := range p1.Inventory {
			fmt.Println(p1.Gold)
			fmt.Printf("%s: %d \n", key, value)

		}
		fmt.Println("\nQue voulez-vous faire ?")
		fmt.Println("1. Sélectionner une potion de soin")
		fmt.Println("2. Sélectionner une potion de poison")
		fmt.Println("3. Sélectionner un sort")
		fmt.Println("4. Retourner en arrière")

		var input int
		fmt.Scan(&input)

		switch input {
		case 1:
			if count, ok := p1.Inventory["Potion de soin"]; ok && count > 0 {
				TakePot(p1)
			} else {
				fmt.Println("Vous n'avez pas de Potion de soin dans votre inventaire.")
			}
		case 2:
			if count, ok := p1.Inventory["Potion de poison"]; ok && count > 0 {
				poisonPot(p1)
			} else {
				fmt.Println("Vous n'avez pas de Potion de poison dans votre inventaire.")
			}
		case 3:
			if count, ok := p1.Inventory["Boule de feu"]; ok && count > 0 {
				SpellBook(p1)
			}
		case 4:
			return
		default:
			fmt.Println("Choix invalide.")
		}
	}
}

func (p1 *Personnage) Dead() {
	if p1.Hp < 0 {
		fmt.Printf("Vous êtes mort ! \n")
		p1.Hp = p1.HpMax / 2
		fmt.Printf("Vous avez été ressuscité avec %d PV", p1.Hp)
	}

}
func (p1 *Personnage) LimiteInventory() bool {
	totalQuantity := 0
	for _, count := range p1.Inventory {
		totalQuantity += count
	}
	return totalQuantity < 10
}
