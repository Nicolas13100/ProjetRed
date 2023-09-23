package game

import (
	"fmt"
	"math/rand"
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
		tc := cases.Title(language.English)
		Name = tc.String(strings.ToLower(Name))
		fmt.Printf("Votre nom sera %s, est-ce correct ? (1.Oui/2.Non) ", Name)
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
				Spells:       []Spell{fireball, iceBlast},
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
				x:          rand.Intn(gridSize),
				y:          rand.Intn(gridSize),
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
				Spells:       []Spell{fireball, iceBlast},
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
				x:          rand.Intn(gridSize),
				y:          rand.Intn(gridSize),
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
				Spells:       []Spell{fireball, iceBlast},
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
				x:          rand.Intn(gridSize),
				y:          rand.Intn(gridSize),
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
				Gold:         9999999,
				InventoryCap: 99999999,
				Spells:       []Spell{fireball, iceBlast},
				Inventory: map[string]int{
					"Potion de soin":   99,
					"Potion de poison": 99,
					"Potion de mana":   99,
					"Fourrure de Loup": 99,
					"Peau de Troll":    99,
					"Cuir de Sanglier": 99,
					"Plume de Corbeau": 99,
				},
				Mana:       999,
				ManaMax:    999,
				Atk:        999,
				Defense:    999,
				Initiative: 999,
				x:          rand.Intn(gridSize),
				y:          rand.Intn(gridSize),
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
		fmt.Println("Sorts", P1.Spells)
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
		fmt.Printf("Vous avez perdu connaissance ! \n")
		fmt.Println("Vous vous retrouve face a une ombre, impossible de savoir de quoi il s'agit")
		fmt.Println("L'ombre vous propose un marché, 5% de vos pièce d'Or en enchange de votre resurection en pleine santé.")
		if askYesNo("Acceptez-vous ?") && p1.Gold > 0 {
			p1.Hp = p1.HpMax
			p1.Mana = p1.ManaMax
			p1.Gold -= p1.Gold * 5 / 100
			fmt.Printf("Vous avez été ressuscité avec %d PV et %d Mana\n", p1.Hp, p1.Mana)
			fmt.Printf("Vous avez été ressuscité avec %d pièce d'Or\n", p1.Gold)
			waitForUserInput("Entrer 0 pour continuer ...")
		} else if askYesNo("Acceptez-vous ?") && p1.Gold < 0 {
			p1.Hp = p1.HpMax / 2
			p1.Mana = p1.ManaMax / 2
			fmt.Printf("Vous n'aviez pas de pièce d'Or")
			fmt.Printf("Vous avez été ressuscité avec %d PV et %d Mana\n", p1.Hp, p1.Mana)
			waitForUserInput("Entrer 0 pour continuer ...")
		} else {
			p1.Hp = p1.HpMax / 2
			p1.Mana = p1.ManaMax / 2
			fmt.Printf("Vous avez été ressuscité avec %d PV et %d Mana\n", p1.Hp, p1.Mana)
			waitForUserInput("Entrer 0 pour continuer ...")
		}
	}

}
