package game

import (
	"fmt"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func CharCreation() *Personnage {
	equipement := Equipment{
		Head:            false,
		Body:            false,
		Leg:             false,
		HPBonus:         0,
		AtkBonus:        0,
		DefBonus:        0,
		InitiativeBonus: 0,
	}
	var Name string
	var Race string
	for {
		ClearConsole()
		fmt.Println("Création de votre personnage :")
		fmt.Println("Seulement des lettres de l'alphabet latin et 10 lettres max")
		// Demander à l'utilisateur de choisir son nom
		fmt.Print("Entrez votre nom : ")

		fmt.Scan(&Name)
		maxLetters := 10
		if !IsValidName(Name, maxLetters) {
			fmt.Printf("Le nom doit contenir uniquement des lettres et avoir au maximum %d caractères.\n", maxLetters)
			continue
		}
		// Vérifier si le nom contient uniquement des lettres
		if !OnlyLetters(Name) {
			fmt.Println("Le nom doit contenir uniquement des lettres.")
			continue
		}

		// Confirm the name
		fmt.Printf("Votre nom est %s, est-ce correct ? (1.Oui/2.Non) ", Name)
		var confirmation int
		fmt.Scan(&confirmation)

		if confirmation == 1 {
			// Create a Title case converter
			tc := cases.Title(language.English)

			// Convert the name to Title case
			Name = tc.String(strings.ToLower(Name))
			break
		} else if confirmation == 2 {
			// If user chooses "no", restart the loop
			continue
		}
	}

	for {
		ClearConsole()
		// Demander à l'utilisateur de choisir sa race (vous pouvez adapter cette partie selon vos besoins)
		fmt.Print("Choisissez votre race : \n")
		fmt.Print("Humain : Vous commencez avec 100 PV Max\n")
		fmt.Print("Elfe : Vous commencez avec 80 PV Max\n")
		fmt.Println("Nain : Vous commencez avec 120 PV Max")
		fmt.Scan(&Race)
		switch Race {
		case "Humain":
			P1 := &Personnage{
				Name:         Name,
				Race:         Race,
				Equipement:   equipement,
				Level:        1,
				Xp:           0,
				XpMax:        100,
				HpMax:        100,
				Hp:           50,
				Gold:         100,
				InventoryCap: 10,
				Skills:       []string{"coup de poing"},
				Inventory: map[string]int{
					"Potion de soin":   3,
					"Potion de poison": 3,
					"Potion de mana":   3,
				},
				Mana:       40,
				ManaMax:    100,
				Atk:        5,
				Defense:    0,
				Initiative: 10,
			}
			return P1

		case "Elfe":
			P1 := &Personnage{
				Name:         Name,
				Race:         Race,
				Equipement:   equipement,
				Level:        1,
				Xp:           0,
				XpMax:        100,
				HpMax:        80,
				Hp:           40,
				Gold:         100,
				InventoryCap: 10,
				Skills:       []string{"coup de poing"},
				Inventory: map[string]int{
					"Potion de soin":   3,
					"Potion de poison": 3,
					"Potion de mana":   3,
				},
				Mana:       50,
				ManaMax:    100,
				Atk:        5,
				Defense:    0,
				Initiative: 20,
			}
			return P1

		case "Nain":
			P1 := &Personnage{
				Name:         Name,
				Race:         Race,
				Equipement:   equipement,
				Level:        1,
				Xp:           0,
				XpMax:        100,
				HpMax:        120,
				Hp:           60,
				Gold:         100,
				InventoryCap: 10,
				Skills:       []string{"coup de poing"},
				Inventory: map[string]int{
					"Potion de soin":   3,
					"Potion de poison": 3,
					"Potion de mana":   3,
				},
				Mana:       40,
				ManaMax:    90,
				Atk:        5,
				Defense:    0,
				Initiative: 5,
			}
			fmt.Printf("Bienvenue à toi : %s\n", Name)
			return P1

		case "Dieu":
			P1 := &Personnage{
				Name:         Name,
				Race:         Race,
				Equipement:   equipement,
				Level:        99,
				HpMax:        999999,
				Hp:           999999,
				Gold:         99999999,
				InventoryCap: 99999999,
				Skills:       []string{"coup de poing"},
				Inventory: map[string]int{
					"Potion de soin":   99,
					"Potion de poison": 99,
					"Potion de mana":   99,
				},
				Mana:       999,
				ManaMax:    999,
				Atk:        999,
				Defense:    999,
				Initiative: 999,
			}
			fmt.Printf("Bienvenue à toi : %s\n", Name)
			return P1
		default:
			fmt.Printf("Race inconnue : %s\n", Race)
			continue
		}
	}
}

func (P1 Personnage) DisplayInfo() {
	// Affichage des informations du Personnage p1

	for {
		ClearConsole()
		fmt.Println("Nom:", P1.Name)
		fmt.Println("Classe:", P1.Race)
		fmt.Println("Niveau:", P1.Level)
		fmt.Println("Points de vie maximum:", P1.HpMax+P1.Equipement.HPBonus)
		fmt.Println("Points de vie actuels:", P1.Hp+P1.Equipement.HPBonus)
		fmt.Println("Cash:", P1.Gold)
		fmt.Println("Sorts", P1.Skills)
		// Affichage de l'inventaire
		fmt.Println("\nType 0 to come back to main menu")

		// Read user input
		var input int

		fmt.Scan(&input)

		// Check if the user wants to return to the main menu
		if input == 0 {
			break
		}
	}
}

func (p1 *Personnage) Dead() {
	if p1.Hp <= 0 {
		fmt.Printf("Vous êtes mort ! \n")
		p1.Hp = p1.HpMax / 2
		fmt.Printf("Vous avez été ressuscité avec %d PV\n", p1.Hp)
	}

}
