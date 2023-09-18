package game

import (
	"fmt"
	"strings"
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

func CharCreation() *Personnage {
	var name string
	var race string
	fmt.Println("Création de votre personnage :")

	// Demander à l'utilisateur de choisir son nom
	fmt.Print("Entrez votre nom : ")
	fmt.Scan(&name)

	// Vérifier si le nom contient uniquement des lettres
	if !OnlyLetters(name) {
		fmt.Println("Le nom doit contenir uniquement des lettres.")
		return nil
	}

	// Mettre la première lettre en majuscule et le reste en minuscule
	name = strings.Title(strings.ToLower(name))

	// Demander à l'utilisateur de choisir sa race (vous pouvez adapter cette partie selon vos besoins)
	fmt.Print("Choisissez votre race : \n")
	fmt.Print("Humain : Vous commencez avec 100 PV Max\n")
	fmt.Print("Elfe : Vous commencez avec 80 PV Max\n")
	fmt.Println("Nain : Vous commencez avec 120 PV Max")
	fmt.Scan(&race)

	switch race {
	case "Humain":
		p1 := &Personnage{
			Name:   name,
			Race:   race,
			Level:  1,
			HpMax:  100,
			Hp:     50,
			Gold:   100,
			Skills: []string{"coup de poing"},
			Inventory: map[string]int{
				"Potion de soin":   3,
				"Potion de poison": 3,
			},
		}
		return p1

	case "Elfe":
		p1 := &Personnage{
			Name:   name,
			Race:   race,
			Level:  1,
			HpMax:  80,
			Hp:     40,
			Gold:   100,
			Skills: []string{"coup de poing"},
			Inventory: map[string]int{
				"Potion de soin":   3,
				"Potion de poison": 3,
			},
		}
		return p1

	case "Nain":
		p1 := &Personnage{
			Name:   name,
			Race:   race,
			Level:  1,
			HpMax:  120,
			Hp:     60,
			Gold:   100,
			Skills: []string{"coup de poing"},
			Inventory: map[string]int{
				"Potion de soin":   3,
				"Potion de poison": 3,
			},
		}
		fmt.Printf("Bienvenue à toi : %s\n", name)
		return p1
	}
	return nil
}

func OnlyLetters(s string) bool {
	for _, char := range s {
		if !('a' <= char && char <= 'z' || 'A' <= char && char <= 'Z') {
			return false
		}
	}
	return true
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
		fmt.Println("Sorts", p1.Skills)
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
		fmt.Println(p1.Gold)
		for key, value := range p1.Inventory {
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
			if count, ok := p1.Inventory["Livre de sort : Boule de feu"]; ok && count > 0 {
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

func (p1 *Personnage) RemoveZeroValueItems() {
	for key, value := range p1.Inventory {
		if value == 0 {
			delete(p1.Inventory, key)
		}
	}
}
